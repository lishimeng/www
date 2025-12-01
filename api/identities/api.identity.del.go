package identities

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/x/container"
)

func apiDelIdentity(ctx server.Context) {
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

	err = identityManager.Unbind(req.IdentityCode, req.UserCode)
	if err != nil {
		log.Info("解绑登录凭证失败: %v", err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "success"
	ctx.Json(resp)
}
