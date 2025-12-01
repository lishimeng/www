package projects

import "github.com/lishimeng/app-starter/server"

func Router(root server.Router) {
	root.Get("/", apiGetSupportProjects)
}
