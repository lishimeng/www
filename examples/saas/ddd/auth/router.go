package auth

import (
	"github.com/lishimeng/app-starter/server"
)

func Router(root server.Router) {
	root.Get("/captcha", getCaptchaApi)
	root.Post("/login", loginApi)
}
