package menus

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/menu"
	"github.com/lishimeng/x/container"
)

type respMenuOption struct {
	app.ResponseWrapper
	Data []*dto.MenuOption `json:"data"`
}

func getMenuOptionsApi(ctx server.Context) {
	var resp respMenuOption
	var flat []*dto.MenuOption
	var err error

	// 获取menuGroup参数，如果未指定则使用当前程序的MenuCategory
	menuGroupParam := ctx.C.URLParamDefault("menuGroup", "")
	var routeOptMgr www.RouteOperationManager
	_ = container.Get(&routeOptMgr)
	list, err := routeOptMgr.GetRouteOptions(menuGroupParam)

	if err != nil {
		_ = log.Error("getMenuOptionsApi failed %s ", err)
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	for _, item := range list {
		flat = append(flat, &item)
	}

	// 构建树形结构，使用buildTree函数
	// 前端需要树形结构来正确显示菜单名称和层级关系
	treeData := menu.BuildTree(flat)
	if treeData == nil {
		// 确保返回空数组而不是 nil
		resp.Data = []*dto.MenuOption{}
	} else {
		resp.Data = treeData
	}

	resp.Code = http.StatusOK
	ctx.Json(resp)
}
