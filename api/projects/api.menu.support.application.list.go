package projects

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www/dto"
)

func apiGetSupportProjects(ctx server.Context) {
	var resp app.ResponseWrapper

	var list []any
	list = append(list, dto.Project{Name: "admin", MenuCategory: "m_admin"})
	list = append(list, dto.Project{Name: "saas", MenuCategory: "m_saas"})
	list = append(list, dto.Project{Name: "manager", MenuCategory: "m_manager"})

	resp.Code = http.StatusOK
	resp.Message = "success"
	resp.Data = list
	ctx.Json(resp)
}
