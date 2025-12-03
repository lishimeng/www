package domain

import (
	"github.com/lishimeng/www"
	"github.com/lishimeng/x/container"
)

type handler struct {
}

func InitManagers() {
	h := handler{}
	// todo: Project的权限管理
	identity := www.NewDbIdentityManager(h)
	container.Add(&identity)
	route := www.NewDbRouteManager(h)
	container.Add(&route)
	routeOpt := www.NewDbRouteOperationManager(h)
	container.Add(&routeOpt)
	prem := www.NewPermManager(h)
	container.Add(&prem)
	premOpt := www.NewDbPermOperationManager(h)
	container.Add(&premOpt)
	role := www.NewRoleManager(h)
	container.Add(&role)
	roleOpt := www.NewRoleOperationManager(h)
	container.Add(&roleOpt)
	auth := www.NewDbAuthManager(h)
	container.Add(&auth)
	authOpt := www.NewDbAuthOperationManager(h)
	container.Add(&authOpt)
	saas := www.NewDbSaasManager(h)
	container.Add(&saas)
	saasOpt := www.NewDbSaasOperationManager(h)
	container.Add(&saasOpt)
}
