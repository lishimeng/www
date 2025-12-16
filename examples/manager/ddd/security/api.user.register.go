package security

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
	"github.com/lishimeng/x/container"
	"net/http"
)

type reqRegister struct {
	IdentityCode string           `json:"identityCode"`
	Password     string           `json:"password"` // todo: encode ?
	IdentityType *int             `json:"identityType"`
	Platform     def.UserPlatform `json:"platform"`
}

func registerApi(ctx server.Context) {
	var req reqRegister
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
		resp.Message = "identity can not be null"
		ctx.Json(resp)
		return
	}

	if req.Platform == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "platform can not be null"
		ctx.Json(resp)
		return
	}

	var identityManager www.IdentityManager
	_ = container.Get(&identityManager)

	// 确定身份类型：优先使用请求中的类型，否则自动判断
	var identityType def.IdentityType
	if req.IdentityType != nil {
		identityType = def.IdentityType(*req.IdentityType)
	} else {
		identityType = identityManager.DetectIdentityType(req.IdentityCode)
	}

	_, user, err := CreateSecurityUser(req.IdentityCode, identityType, req.Platform)
	if err != nil {
		log.Info("创建用户失败: %v", err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	// 设置密码
	err = UpdatePsw(user.Code, req.Password, func(_ usersTable.AuthSecurityInfo) bool {
		return true
	})
	if err != nil {
		log.Info("设置密码失败: %v", err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Data = user
	ctx.Json(resp)
}
