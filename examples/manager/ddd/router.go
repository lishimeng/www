package ddd

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www/api/identities"
	"github.com/lishimeng/www/api/permissions"
	"github.com/lishimeng/www/api/roles"
	"github.com/lishimeng/www/api/sdk"
	"github.com/lishimeng/www/api/users"
	"github.com/lishimeng/www/examples/manager/ddd/file"
	"github.com/lishimeng/www/examples/manager/ddd/profile"
	"github.com/lishimeng/www/examples/manager/ddd/security"
)

func Route(app server.Router) {
	router(app.Path("/api"))
}

func router(root server.Router) {
	identities.Router(root.Path("/identities"))
	sdk.Router(root.Path("/menus"))
	permissions.Router(root.Path("/permissions"))
	roles.Router(root.Path("/roles"))
	users.Router(root.Path("/users"))
	security.Router(root.Path("/users"))
	profile.Router(root.Path("/profile"))

	file.Router(root.Path("/v1/files"))
}
