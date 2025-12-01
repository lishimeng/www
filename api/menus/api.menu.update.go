package menus

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/x/container"
)

func updateMenuApi(ctx server.Context) {
	var req dto.ReqMenuForm
	var resp app.Response

	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		ctx.Json(resp)
		return
	}

	id := ctx.C.Params().GetString("id")
	req.Id = id

	var routeOptMng www.RouteOperationManager
	_ = container.Get(&routeOptMng)
	err = routeOptMng.EditRoute(&req)

	if err != nil {
		resp.Message = err.Error()
		resp.Code = http.StatusBadRequest
		ctx.Json(resp)
		return
	}
	resp.Code = http.StatusOK
	ctx.Json(resp)
}
