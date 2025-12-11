package identities

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/x/container"
	"net/http"
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

	var identityManager www.IdentityManager
	_ = container.Get(&identityManager)

	// 确定身份类型：优先使用请求中的类型，否则自动判断
	var identityType def.IdentityType
	if req.IdentityType != nil {
		identityType = def.IdentityType(*req.IdentityType)
	} else {
		identityType = identityManager.DetectIdentityType(req.IdentityCode)
	}

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
