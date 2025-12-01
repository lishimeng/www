package menus

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/x/container"
)

func deleteMenuApi(ctx server.Context) {
	var resp app.Response
	// 前端传递的是string类型的ID，需要先获取string再转换为int
	id := ctx.C.Params().GetIntDefault("id", 0)
	if id == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "id parameter is required"
		ctx.Json(resp)
		return
	}

	var routeOptMgr www.RouteOperationManager
	_ = container.Get(&routeOptMgr)
	err := routeOptMgr.DelRoute(id)

	if err != nil {
		log.Info(err.Error())
		resp.Code = http.StatusNotFound
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	resp.Code = http.StatusOK
	ctx.Json(resp)
}
