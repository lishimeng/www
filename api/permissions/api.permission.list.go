package permissions

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
	"github.com/lishimeng/x/container"
)

func apiPermissionList(ctx server.Context) {
	var resp app.ResponseWrapper

	var permOptMgr www.PermOperationManager
	_ = container.Get(&permOptMgr)
	list := permOptMgr.AllPerms()

	resp.Code = http.StatusOK
	resp.Data = list
	ctx.Json(resp)
}
