package www

import "github.com/lishimeng/www/dto"

type DbRoleManager struct {
	domain RoleDomainHandler
}

func NewRoleManager(domain RoleDomainHandler) RoleManager {
	var h RoleManager
	h = &DbRoleManager{domain: domain}
	return h
}

func (rm *DbRoleManager) UserRoles(userCode string) (role []dto.UserRole, err error) {

	role, err = rm.domain.GetUserRoleList(userCode)
	return
}
