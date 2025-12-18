package domains

import (
	"errors"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/examples/internal/db/auth/permissionsTable"
)

func (h handler) GetUserRoleList(userCode string) (role []dto.UserRole, err error) {
	var models []permissionsTable.AuthUserRoles
	_, err = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthUserRoles)).
		Filter("SaasUserCode", userCode).
		OrderBy("RoleCategoryCode", "Id").
		All(&models)
	if err != nil {
		return
	}
	for _, one := range models {
		var oneDto dto.UserRole
		one.Transform(&oneDto)
		role = append(role, oneDto)
	}
	return
}

func (h handler) GetAllRole() (list []dto.Role, err error) {
	var model []permissionsTable.AuthRoles
	_, err = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthRoles)).
		OrderBy("-RoleCategoryCode", "Id").
		All(&model)
	if err != nil {
		return
	}
	for _, one := range model {
		var oneDto dto.Role
		one.Transform(&oneDto)
		list = append(list, oneDto)
	}
	return
}

func (h handler) CreateRole(role *dto.Role) (err error) {
	var model permissionsTable.AuthRoles
	model.Convert(*role)
	_, err = app.GetOrm().Context.Insert(&model)
	return
}

func (h handler) UpdateRole(role *dto.Role) (err error) {
	var model permissionsTable.AuthRoles
	model.Convert(*role)
	_, err = app.GetOrm().Context.Update(&model)
	return
}

func (h handler) DelRole(roleCode string) (err error) {
	_, err = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthRoles)).
		Filter("Code", roleCode).
		Delete()
	return err
}

func (h handler) GetRoleByCode(roleCode string) (role dto.Role, err error) {
	var model permissionsTable.AuthRoles
	model.Code = roleCode
	if err = app.GetOrm().Context.Read(&model, "Code"); err != nil {
		return
	}
	model.Transform(&role)
	return
}

func (h handler) CreateUserRole(saasUserCode, roleCode string) (ur dto.UserRole, err error) {
	// 先检查重复
	repeated := app.GetOrm().Context.QueryTable(new(permissionsTable.AuthUserRoles)).
		Filter("SaasUserCode", saasUserCode).
		Filter("RoleCode", roleCode).
		Exist()
	if repeated {
		err = errors.New("用户角色绑定重复")
		return
	}
	var model permissionsTable.AuthUserRoles
	model.SaasUserCode = saasUserCode
	model.RoleCode = roleCode
	model.IsGlobal = 1
	// todo: 暂时全global，后续升级; 与 role的 global字段一致？
	// RoleCategoryCode 与role表一致？
	_, err = app.GetOrm().Context.Insert(&model)
	return
}

func (h handler) DeleteUserRole(saasUserCode, roleCode string) (err error) {
	_, err = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthUserRoles)).
		Filter("SaasUserCode", saasUserCode).
		Filter("RoleCode", roleCode).
		Delete()
	return
}

func (h handler) GetUserRole(filterUserCode, filterRoleCode string) (list []dto.UserRole, err error) {
	var models []permissionsTable.AuthUserRoles
	q := app.GetOrm().Context.QueryTable(new(permissionsTable.AuthUserRoles))
	if len(filterUserCode) > 0 {
		q = q.Filter("SaasUserCode", filterUserCode)
	}
	if len(filterRoleCode) > 0 {
		q = q.Filter("RoleCode", filterRoleCode)
	}
	_, err = q.OrderBy("RoleCategoryCode", "Id").All(&models)
	if err != nil {
		return
	}
	for _, one := range models {
		var oneDto dto.UserRole
		one.Transform(&oneDto)
		list = append(list, oneDto)
	}
	return
}
