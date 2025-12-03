package domain

import (
	"errors"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/examples/internal/db/auth/permissionsTable"
)

func (h handler) GetUserRoleList(userCode string) (role []dto.UserRole, err error) {
	var models []permissionsTable.AuthUserRoles
	_, err = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthUserRoles)).
		Filter("UserCode", userCode).
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
	model.Transform(role)
	_, err = app.GetOrm().Context.Insert(&model)
	return
}

func (h handler) UpdateRole(role *dto.Role) (err error) {
	var model permissionsTable.AuthRoles
	model.Transform(role)
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

func (h handler) CreateUserRole(userCode, roleCode string) (ur dto.UserRole, err error) {
	// 先检查重复
	repeated := app.GetOrm().Context.QueryTable(new(permissionsTable.AuthUserRoles)).
		Filter("UserCode", userCode).
		Filter("RoleCode", roleCode).
		Exist()
	if repeated {
		err = errors.New("用户角色绑定重复")
		return
	}
	var model permissionsTable.AuthUserRoles
	model.UserCode = userCode
	model.RoleCode = roleCode
	model.IsGlobal = 1
	// todo: 暂时全global，后续升级
	_, err = app.GetOrm().Context.Insert(&model)
	return
}

func (h handler) DeleteUserRole(userCode, roleCode string) (err error) {
	_, err = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthRolePermissions)).
		Filter("UserCode", userCode).
		Filter("RoleCode", roleCode).
		Delete()
	return
}

func (h handler) GetUserRole(filterUserCode, filterRoleCode string) (list []dto.UserRole, err error) {
	var models []permissionsTable.AuthUserRoles
	_, err = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthUserRoles)).
		Filter("UserCode", filterUserCode).
		Filter("RoleCode", filterRoleCode).
		All(&models)
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
