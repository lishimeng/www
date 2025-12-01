package menus

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
)

func Router(root server.Router) {
	root.Get("/routes", www.WithAuth(GetRoutesApi)...)
	root.Get("/", www.WithAuth(getMenuListApi)...)
	root.Get("/options", www.WithAuth(getMenuOptionsApi)...)
	root.Get("/{id}/form", www.WithAuth(getMenuFormApi)...)
	root.Post("/", www.WithAuth(createMenuApi)...)
	root.Put("/{id}", www.WithAuth(updateMenuApi)...)
	root.Delete("/{id}", www.WithAuth(deleteMenuApi)...)
}
