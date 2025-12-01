package permissions

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/x/container"
)

type FormRolePermission struct {
	RoleCode       string `json:"roleCode,omitempty"`
	PermissionCode string `json:"permissionCode,omitempty"`
}

func apiRoleBindPermission(ctx server.Context) {
	var req FormRolePermission
	var err error
	var resp app.ResponseWrapper

	err = ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	if len(req.RoleCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "role code can not be null"
		ctx.Json(resp)
		return
	}
	if len(req.PermissionCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "permission code can not be null"
		ctx.Json(resp)
		return
	}

	var permOptMgr www.PermOperationManager
	_ = container.Get(&permOptMgr)
	err = permOptMgr.BindRole(req.RoleCode, req.PermissionCode)

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

func apiRoleUnbindPermission(ctx server.Context) {
	var req FormRolePermission
	var err error
	var resp app.ResponseWrapper

	err = ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	if len(req.RoleCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "role code can not be null"
		ctx.Json(resp)
		return
	}
	if len(req.PermissionCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "permission code can not be null"
		ctx.Json(resp)
		return
	}

	var permOptMgr www.PermOperationManager
	_ = container.Get(&permOptMgr)
	err = permOptMgr.UnbindRole(req.RoleCode, req.PermissionCode)
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
