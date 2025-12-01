package identities

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/x/container"
)

func apiAddIdentity(ctx server.Context) {
	var req IdentityForm
	var err error
	var resp app.ResponseWrapper

	err = ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	if len(req.IdentityCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "identity code can not be null"
		ctx.Json(resp)
		return
	}

	if len(req.UserCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "user code can not be null"
		ctx.Json(resp)
		return
	}

	// 确定身份类型：优先使用请求中的类型，否则自动判断
	var identityType def.IdentityType
	if req.IdentityType != nil {
		identityType = def.IdentityType(*req.IdentityType)
	} else {
		identityType = detectIdentityType(req.IdentityCode)
	}

	var identityManager www.IdentityManager
	_ = container.Get(&identityManager)

	err = identityManager.Bind(req.IdentityCode, req.UserCode, identityType)
	if err != nil {
		log.Info("绑定登录凭证失败: %v", err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "success"
	ctx.Json(resp)
}

// detectIdentityType 根据 identityCode 的格式自动判断身份类型
func detectIdentityType(identityCode string) def.IdentityType {
	// 判断是否为邮箱（包含 @ 符号）
	if strings.Contains(identityCode, "@") {
		return def.LoginEmail
	}

	// 判断是否为手机号（11位数字，以1开头）
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if phoneRegex.MatchString(identityCode) {
		return def.LoginPhone
	}

	// 判断是否为微信相关（包含特定前缀或格式）
	if strings.HasPrefix(identityCode, "wx_") || strings.HasPrefix(identityCode, "unionid_") {
		return def.LoginWechatUnionId
	}
	if strings.HasPrefix(identityCode, "openid_") {
		return def.LoginWechatOpenId
	}

	// 默认使用登录码
	return def.LoginCode
}
