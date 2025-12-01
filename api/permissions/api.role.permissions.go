package permissions

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
	"github.com/lishimeng/x/container"
)

func apiRolePermissions(ctx server.Context) {
	var resp app.ResponseWrapper
	roleCode := ctx.C.Params().GetStringDefault("roleCode", "")

	if len(roleCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "role code is required"
		ctx.Json(resp)
		return
	}

	var permOptMgr www.PermOperationManager
	_ = container.Get(&permOptMgr)
	list := permOptMgr.GetRolePerms(roleCode)

	resp.Code = http.StatusOK
	resp.Data = list
	ctx.Json(resp)
}
