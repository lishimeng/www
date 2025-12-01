package identities

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
)

func Router(root server.Router) {
	root.Get("/", www.WithAuth(apiIdentityList)...)
	root.Post("/", www.WithAuth(apiAddIdentity)...)   // 创建一个登录凭证
	root.Delete("/", www.WithAuth(apiDelIdentity)...) // 删除一个登录凭证
}
