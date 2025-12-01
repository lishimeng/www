package www

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/www/dto"
)

type DbPermOperationManager struct {
	domain DomainHandler
}

func NewDbPermOperationManager(db DomainHandler) PermOperationManager {
	var h PermOperationManager
	h = &DbPermOperationManager{domain: db}
	return h
}

func (h *DbPermOperationManager) CreatePerm(dto *dto.Permission) (err error) {
	err = app.GetOrm().Transaction(func(tx persistence.TxContext) (e error) {
		e = h.domain.TxCreatePerm(tx, dto)
		return
	})
	return
}
func (h *DbPermOperationManager) EditPerm(code string, dto *dto.Permission) (err error) {
	dto.Code = code
	err = app.GetOrm().Transaction(func(tx persistence.TxContext) (e error) {
		e = h.domain.TxUpdatePerm(tx, dto)
		return
	})
	return
}

func (h *DbPermOperationManager) AllPerms() (perms []dto.Permission) {
	perms = h.domain.GetAllPerms()
	return
}

func (h *DbPermOperationManager) DelPerm(code string) (err error) {
	err = app.GetOrm().Transaction(func(tx persistence.TxContext) (e error) {
		e = h.domain.TxDeletePerm(tx, code)
		return
	})
	return
}

func (h *DbPermOperationManager) One(code string) (dto dto.Permission, err error) {
	dto, err = h.domain.GetPermByCode(code)
	return
}

func (h *DbPermOperationManager) BindRole(roleCode, permCome string) (err error) {
	p := dto.AuthRolePerm{
		RoleCode:       roleCode,
		PermissionCode: permCome,
	}
	err = app.GetOrm().Transaction(func(tx persistence.TxContext) (e error) {

		e = h.domain.TxCreateRolePermNoChk(tx, &p)
		if e != nil {
			return
		}
		return
	})

	return
}
func (h *DbPermOperationManager) UnbindRole(roleCode, permCome string) (err error) {
	p := dto.AuthRolePerm{
		RoleCode:       roleCode,
		PermissionCode: permCome,
	}
	err = app.GetOrm().Transaction(func(tx persistence.TxContext) (e error) {
		e = h.domain.TxDeleteRolePermNoChk(tx, &p)
		return
	})
	return
}

func (h *DbPermOperationManager) GetRolePerms(roleCode string) (perms []dto.Permission) {
	perms = h.domain.GetPermsByRole(roleCode)
	return
}
