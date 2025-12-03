package permissionsTable

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/www/dto"
)

// 权限列表结构: 应用:组织:用户:权限列表
// 建议只在saas_user/admin_user受role规则管理

type AuthRolesCategory struct { // 类型标签
	app.Pk
	Code string `orm:"column(code)"` // 编号
	Name string `orm:"column(name)"` // 名称
	app.TableChangeInfo
}

type AuthRoles struct { // 角色
	app.Pk
	Code             string `orm:"column(code);unique"`
	Name             string `orm:"column(name)"`               // 名称
	RoleCategoryCode string `orm:"column(role_category_code)"` // 角色类型(管理角色树)
	ProjectCode      string `orm:"column(project_code);null"`  // 系统编号(系统内专用权限)
	IsGlobal         int    `orm:"column(is_global)"`          // 全局权限(与系统内权限互斥)
	app.TableChangeInfo
}

func (ar *AuthRoles) IsGlobalRole() bool {
	return ar.IsGlobal == 1
}

func (ar *AuthRoles) Transform(dst *dto.Role) {
	dst.Id = ar.Id
	dst.Code = ar.Code
	dst.Name = ar.Name
	dst.RoleCategoryCode = ar.RoleCategoryCode
	dst.ProjectCode = ar.ProjectCode
	dst.RoleCategoryCode = ar.RoleCategoryCode
	dst.IsGlobal = ar.IsGlobal == 1

}

type AuthUserRoles struct { // 用户权限列表, 使用user和项目(project/global)确定权限范围
	app.Pk
	UserCode         string `orm:"column(user_code)"`
	RoleCategoryCode string `orm:"column(role_category_code)"` // 角色类型/冗余
	RoleCode         string `orm:"column(role_code)"`
	ProjectCode      string `orm:"column(project_code);null"` // 冗余
	IsGlobal         int    `orm:"column(is_global)"`         // 冗余
	app.TableChangeInfo
}

func (ur *AuthUserRoles) Transform(dst *dto.UserRole) {
	dst.Id = ur.Id
	dst.RoleCode = ur.RoleCode
	dst.UserCode = ur.UserCode
	dst.ProjectCode = ur.ProjectCode
	dst.RoleCategoryCode = ur.RoleCategoryCode
	dst.IsGlobal = ur.IsGlobal == 1
}
