package auth

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
	"github.com/lishimeng/www/examples/saas/storage"
	"github.com/lishimeng/www/examples/saas/viewToken"
	"github.com/lishimeng/x/container"
	"net/http"
	"time"
)

type formLogin struct {
	Username    string `form:"username"`
	Password    string `form:"password"`
	CaptchaKey  string `form:"captchaKey"`
	CaptchaCode string `form:"captchaCode"`
}

func loginApi(ctx server.Context) {
	var req formLogin
	var err error
	var resp app.ResponseWrapper

	err = ctx.C.ReadForm(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	// 查询保存的answer
	var answer any
	answer, err = storage.Get(captchaPrefix + req.CaptchaKey)
	if err != nil {
		log.Info("get captcha error")
		resp.Code = http.StatusInternalServerError
		resp.Message = "服务器异常"
		ctx.Json(resp)
		return
	}
	if ans, ok := answer.(string); !ok || ans != req.CaptchaCode {
		log.Info("captcha answer error")
		resp.Code = http.StatusNotFound
		resp.Message = "验证码错误"
		ctx.Json(resp)
		return
	}
	_ = storage.Del(captchaPrefix + req.CaptchaKey)

	// 验证码正确，正式进入登录逻辑
	identity := usersTable.AuthIdentity{Code: req.Username}
	err = app.GetOrm().Context.Read(&identity, "Code")
	if err != nil {
		resp.Code = http.StatusNotFound
		resp.Message = "未注册的账号"
		ctx.Json(resp)
		return
	}

	security := usersTable.AuthSecurityInfo{Code: identity.SecurityCode}
	err = app.GetOrm().Context.Read(&security, "Code")
	if err != nil {
		resp.Code = http.StatusNotFound
		resp.Message = "账号状态异常"
		ctx.Json(resp)
		return
	}

	if !security.ValidatePassword(req.Password) {
		resp.Code = http.StatusNotFound
		resp.Message = "密码错误"
		ctx.Json(resp)
		return
	}

	// todo: 可以配置取消platform校验
	if security.Platform != www.SystemScope.Code {
		resp.Code = http.StatusNotFound
		resp.Message = "不是本系统的账号"
		ctx.Json(resp)
		return
	}

	// saas project 独有：查询用户的saas组织
	var saasMgr www.SaasManager
	_ = container.Get(&saasMgr)
	saasUsers, err := saasMgr.UserSaasOrgs(security.Code)
	if err != nil || len(saasUsers) == 0 {
		log.Info(err)
		resp.Code = http.StatusNotFound
		resp.Message = ""
		ctx.Json(resp)
		return
	}

	// todo if len(saasUsers) > 1 {}

	// 校验成功，生成token
	bs, err := viewToken.GenToken(
		saasUsers[0].Code, "password",
		www.SystemScope.Name, time.Minute,
	)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = "服务器异常"
		ctx.Json(resp)
		return
	}
	var dto viewToken.DtoViewAccessToken
	dto.TokenType = "Bearer"
	dto.AccessToken = string(bs)
	dto.RefreshToken = viewToken.GenRefreshToken(dto.AccessToken)
	dto.ExpiresIn = 720 // todo: 单位是秒？需要确认

	resp.Code = http.StatusOK
	resp.Data = dto
	resp.Message = "success"
	ctx.Json(resp)
}
