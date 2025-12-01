package permissions

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/x/container"
)

func apiDelPermission(ctx server.Context) {
	var err error
	var resp app.ResponseWrapper
	permissionCode := ctx.C.Params().GetStringDefault("code", "")

	if len(permissionCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "permission code is required"
		ctx.Json(resp)
		return
	}

	var permOptMgr www.PermOperationManager
	_ = container.Get(&permOptMgr)
	err = permOptMgr.DelPerm(permissionCode)

	if err != nil {
		log.Info(err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "success"
	ctx.Json(resp)
}
