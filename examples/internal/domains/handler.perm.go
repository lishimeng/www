package domains

import (
	"errors"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/examples/internal/db/auth/permissionsTable"
)

func (h handler) GetPermissionList(filterUserCode string, _ string) (list []dto.AuthUserPerm, err error) {
	var views []permissionsTable.AuthUserPermView
	_, err = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthUserPermView)).
		Filter("UserCode", filterUserCode).
		OrderBy("Id").
		All(&views)
	if err != nil {
		return
	}
	for _, one := range views {
		var oneDto dto.AuthUserPerm
		one.Transform(&oneDto)
		list = append(list, oneDto)
	}
	return
}

func (h handler) TxCreatePerm(tx persistence.TxContext, dto *dto.Permission) (err error) {
	var model permissionsTable.AuthPermissions
	model.Convert(*dto)
	_, err = tx.Context.Insert(&model)
	return
}

func (h handler) TxUpdatePerm(tx persistence.TxContext, dto *dto.Permission) (err error) {
	var model permissionsTable.AuthPermissions
	var id int
	model.Code = dto.Code
	if err = tx.Context.Read(&model, "Code"); err != nil {
		return
	}
	id = model.Id
	model.Convert(*dto)
	model.Id = id
	_, err = tx.Context.Update(&model)
	return
}

func (h handler) TxDeletePerm(tx persistence.TxContext, code string) (err error) {
	_, err = tx.Context.QueryTable(new(permissionsTable.AuthPermissions)).
		Filter("Code", code).
		Delete()
	return
}

func (h handler) GetAllPerms() (perms []dto.Permission) {
	var models []permissionsTable.AuthPermissions
	_, _ = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthPermissions)).
		OrderBy("Id").
		All(&models)
	for _, model := range models {
		var one dto.Permission
		model.Transform(&one)
		perms = append(perms, one)
	}
	return
}

func (h handler) GetPermsByRole(roleCode string) (perms []dto.Permission) {
	var models []permissionsTable.AuthRolePermissions
	n, _ := app.GetOrm().Context.QueryTable(new(permissionsTable.AuthRolePermissions)).
		Filter("RoleCode", roleCode).
		All(&models)
	if n == 0 {
		return
	}
	var permCodes []string
	for _, one := range models {
		permCodes = append(permCodes, one.PermissionCode)
	}
	var permModels []permissionsTable.AuthPermissions
	_, _ = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthPermissions)).
		Filter("Code__in", permCodes).OrderBy("Id").All(&permModels)
	for _, one := range permModels {
		var oneDto dto.Permission
		one.Transform(&oneDto)
		perms = append(perms, oneDto)
	}
	return
}

func (h handler) GetPermByCode(code string) (dto dto.Permission, err error) {
	var model permissionsTable.AuthPermissions
	model.Code = code
	if err = app.GetOrm().Context.Read(&model, "Code"); err != nil {
		return
	}
	model.Transform(&dto)
	return
}

func (h handler) TxCreateRolePermNoChk(tx persistence.TxContext, dto *dto.AuthRolePerm) (err error) {
	// 先检查重复
	var model permissionsTable.AuthRolePermissions
	model.Convert(*dto)
	err = tx.Context.Read(&model, "RoleCode", "PermissionCode")
	if err == nil {
		return errors.New("重复绑定角色权限")
	}
	_, err = tx.Context.Insert(&model)
	return
}

func (h handler) TxDeleteRolePermNoChk(tx persistence.TxContext, dto *dto.AuthRolePerm) (err error) {
	_, err = tx.Context.QueryTable(new(permissionsTable.AuthRolePermissions)).
		Filter("RoleCode", dto.RoleCode).
		Filter("PermissionCode", dto.PermissionCode).
		Delete()
	return
}
