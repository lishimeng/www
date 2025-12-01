package permissions

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/x/container"
)

func apiUpdatePermission(ctx server.Context) {
	var err error
	var req dto.Permission
	var resp app.ResponseWrapper

	err = ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	if len(req.Code) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "permission code can not be null"
		ctx.Json(resp)
		return
	}
	var permOptMgr www.PermOperationManager
	_ = container.Get(&permOptMgr)
	err = permOptMgr.EditPerm(req.Code, &req)

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
