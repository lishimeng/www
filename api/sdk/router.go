package sdk

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/api/menus"
)

func Router(root server.Router) {
	root.Get("/routes", www.WithAuth(menus.GetRoutesApi)...)
}
