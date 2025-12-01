package menus

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/x/container"
)

// getMenuFormApi 获取菜单表单数据
func getMenuFormApi(ctx server.Context) {
	var err error
	var resp app.ResponseWrapper

	idStr := ctx.C.Params().GetString("id")
	if idStr == "" {
		resp.Code = http.StatusBadRequest
		resp.Message = "id parameter is required"
		ctx.Json(resp)
		return
	}

	id := dto.ConvertInt(idStr)
	if id == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "invalid id parameter"
		ctx.Json(resp)
		return
	}
	var menuManager www.RouteOperationManager
	_ = container.Get(&menuManager)
	menu, err := menuManager.GetRoute(id)

	if err != nil {
		_ = log.Error("getMenuFormApi failed %s ", err)
		resp.Code = http.StatusNotFound
		resp.Message = "menu not found"
		ctx.Json(resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Data = menu
	ctx.Json(resp)
}
