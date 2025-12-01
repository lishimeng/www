package roles

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www/dto"
)

func apiRoleCategoryList(ctx server.Context) {
	var resp app.ResponseWrapper
	var list []dto.RoleCategory

	// TODO 硬编码
	list = append(list, dto.RoleCategory{
		Id:   1,
		Code: "rc_001",
		Name: "一级权限",
	})
	list = append(list, dto.RoleCategory{
		Id:   2,
		Code: "rc_002",
		Name: "二级权限",
	})
	list = append(list, dto.RoleCategory{
		Id:   3,
		Code: "rc_003",
		Name: "三级权限",
	})
	list = append(list, dto.RoleCategory{
		Id:   4,
		Code: "rc_004",
		Name: "四级权限",
	})
	resp.Code = http.StatusOK
	resp.Data = list
	ctx.Json(resp)
}
