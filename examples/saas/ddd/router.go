package ddd

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www/api/identities"
	"github.com/lishimeng/www/api/sdk"
	"github.com/lishimeng/www/api/users"
	"github.com/lishimeng/www/examples/saas/ddd/auth"
)

func Route(app server.Router) {
	router(app.Path("/api"))
}

func router(root server.Router) {
	identities.Router(root.Path("/identities"))
	sdk.Router(root.Path("/menus"))
	users.Router(root.Path("/users"))
	auth.Router(root.Path("/auth"))
}
