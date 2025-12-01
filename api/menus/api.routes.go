package menus

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/menu"
	"github.com/lishimeng/x/container"
)

func GetRoutesApi(ctx server.Context) {
	var resp app.ResponseWrapper
	var flat []*dto.Menu

	// 获取用户权限列表
	routes := getRouters(ctx)

	for _, item := range routes {
		flat = append(flat, &item)
	}

	resp.Code = http.StatusOK

	builtRoutes := menu.BuildRoutes(menu.BuildTree(flat), "")
	// 确保builtRoutes不是nil，如果是nil或空则使用空数组
	if len(builtRoutes) == 0 {
		resp.Data = []*dto.MenuRoute{} // 返回空数组而不是nil
	} else {
		resp.Data = builtRoutes
	}
	ctx.Json(resp)
}

func getRouters(ctx server.Context) (routers []dto.Menu) {
	uid := auth.GetUid(ctx)
	var menuManager www.RouteManager
	_ = container.Get(&menuManager)
	routers = menuManager.GetUserRouters(uid)

	return
}
