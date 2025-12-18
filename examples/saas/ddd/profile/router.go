package profile

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
)

func Router(root server.Router) {
	root.Get("/profile", www.WithAuth(getProfileApi)...)
	root.Get("/me", www.WithAuth(getMeApi)...)
}
