package profile

import "github.com/lishimeng/app-starter/server"

func Router(root server.Router) {
	root.Post("/{code}", initializeApi)
	root.Get("/{code}", getOneApi)
	root.Patch("/{code}", updateApi)
}
