package menus

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/x/container"
)

func createMenuApi(ctx server.Context) {
	var req dto.ReqMenuForm
	var resp app.Response

	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		ctx.Json(resp)
		return
	}

	var menuOptManager www.RouteOperationManager
	_ = container.Get(&menuOptManager)
	err = menuOptManager.AddRoute(&req)

	if err != nil {
		resp.Message = err.Error()
		resp.Code = http.StatusBadRequest
		ctx.Json(resp)
		return
	}
	resp.Code = http.StatusOK
	ctx.Json(resp)
}
