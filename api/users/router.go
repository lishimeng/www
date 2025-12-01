package users

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
	"github.com/lishimeng/x/container"
)

func getAuthManager() www.AuthManager {
	var m www.AuthManager
	_ = container.Get(&m)
	return m
}

func getAuthOptManager() www.AuthOperationManager {
	var m www.AuthOperationManager
	_ = container.Get(&m)
	return m
}

func Router(root server.Router) {
	root.Get("/saas/list", www.WithAuth(apiSaasUserList)...)
	root.Post("/saas/create", www.WithAuth(apiCreateSaasUser)...)
	root.Get("/me", www.WithAuth(apiMeInfo)...)                        // 登录人
	root.Get("/detail/{securityCode}", www.WithAuth(apiUserDetail)...) // 用户详情（包含登录凭证列表）
}
