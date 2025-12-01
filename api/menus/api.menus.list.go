package menus

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/menu"
	"github.com/lishimeng/x/container"
)

// getMenuListApi 不分页
func getMenuListApi(ctx server.Context) {
	var resp app.ResponseWrapper
	var flat []*dto.Menu

	// 获取关键词搜索参数
	keywords := ctx.C.URLParam("keywords")
	// 获取menuGroup参数，如果未指定则使用当前程序的MenuCategory
	menuGroupParam := ctx.C.URLParamDefault("menuGroup", "")

	var routeOptManager www.RouteOperationManager
	_ = container.Get(&routeOptManager)
	list := routeOptManager.GetProjectRoutes(keywords, menuGroupParam)

	for _, item := range list {
		flat = append(flat, &item)
	}
	resp.Data = menu.BuildTree(flat)
	resp.Code = http.StatusOK
	ctx.Json(resp)
}
