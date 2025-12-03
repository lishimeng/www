package ddd

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www/api/identities"
	"github.com/lishimeng/www/api/menus"
	"github.com/lishimeng/www/api/permissions"
	"github.com/lishimeng/www/api/projects"
	"github.com/lishimeng/www/api/roles"
	"github.com/lishimeng/www/api/users"
)

func Route(app server.Router) {
	router(app.Path("/api"))
}

func router(root server.Router) {
	identities.Router(root.Path("/identities"))
	menus.Router(root.Path("/menus"))
	permissions.Router(root.Path("/permissions"))
	projects.Router(root.Path("/projects"))
	roles.Router(root.Path("/roles"))
	users.Router(root.Path("/users"))
}
