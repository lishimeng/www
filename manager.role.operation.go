package www

import (
	"github.com/lishimeng/www/dto"
)

type DbRoleOperationManager struct {
	domain RoleOperationDomainHandler
}

func NewRoleOperationManager(domain RoleOperationDomainHandler) RoleOperationManager {
	var h RoleOperationManager
	h = &DbRoleOperationManager{domain: domain}
	return h
}

func (h *DbRoleOperationManager) All() (list []dto.Role, err error) {
	list, err = h.domain.GetAllRole()
	return
}

func (h *DbRoleOperationManager) AddRole(role *dto.Role) (err error) {
	err = h.domain.CreateRole(role)
	return
}

func (h *DbRoleOperationManager) DelRole(roleCode string) (err error) {
	err = h.domain.DelRole(roleCode)
	return
}

func (h *DbRoleOperationManager) One(roleCode string) (one dto.Role, err error) {
	one, err = h.domain.GetRoleByCode(roleCode)
	return
}

func (h *DbRoleOperationManager) EditRole(role *dto.Role) (err error) {
	err = h.domain.UpdateRole(role)
	return
}

func (h *DbRoleOperationManager) BindUser(userCode, roleCode string) (ur dto.UserRole, err error) {
	ur, err = h.domain.CreateUserRole(userCode, roleCode)
	return
}

func (h *DbRoleOperationManager) UnbindUser(userCode, roleCode string) (err error) {
	err = h.domain.DeleteUserRole(userCode, roleCode)
	return
}

func (h *DbRoleOperationManager) RoleUsers(roleCode string) (list []dto.UserRole, err error) {
	list, err = h.domain.GetUserRole("", roleCode)
	return
}
