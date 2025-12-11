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

func getSaasManager() www.SaasManager {
	var m www.SaasManager
	_ = container.Get(&m)
	return m
}

func getSaasOptManager() www.SaasOperationManager {
	var m www.SaasOperationManager
	_ = container.Get(&m)
	return m
}

func Router(root server.Router) {
	root.Get("/saas/list", www.WithAuth(apiSaasUserList)...)
	root.Post("/saas/create", www.WithAuth(apiCreateSaasUser)...)
	root.Post("/saas/join", www.WithAuth(apiSaasJoin)...)
	root.Post("/saas/leave", www.WithAuth(apiSaasLeave)...)
	root.Get("/saas/orgs", www.WithAuth(apiSaasList)...)
	root.Get("/saas/orgs/{orgCode}", www.WithAuth(apiSassOne)...)
	root.Get("/saas/my", www.WithAuth(apiMyOrgs)...)
	root.Get("/me", www.WithAuth(apiMeInfo)...)                        // 登录人 todo
	root.Get("/detail/{securityCode}", www.WithAuth(apiUserDetail)...) // 用户详情（包含登录凭证列表）
}
