package permissionsTable

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/www/dto"
)

// AuthPermissions 权限表
type AuthPermissions struct {
	app.Pk
	Code        string `orm:"column(code);unique"`      // 权限编码，唯一标识
	Name        string `orm:"column(name)"`             // 权限名称
	Description string `orm:"column(description);null"` // 权限描述
	Category    string `orm:"column(category);null"`    // 权限分类
	app.TableChangeInfo
}

// AuthRolePermissions 角色权限关联表
type AuthRolePermissions struct {
	app.Pk
	RoleCode       string `orm:"column(role_code)"`       // 角色编码
	PermissionCode string `orm:"column(permission_code)"` // 权限编码
	app.TableChangeInfo
}

func (rp *AuthRolePermissions) Transform(dst *dto.AuthRolePerm) {
	dst.Id = rp.Id
	dst.RoleCode = rp.RoleCode
	dst.PermissionCode = rp.PermissionCode
}

func (rp *AuthRolePermissions) Convert(src dto.AuthRolePerm) {
	rp.Id = src.Id
	rp.RoleCode = src.RoleCode
	rp.PermissionCode = src.PermissionCode
}

func (p *AuthPermissions) Transform(dst *dto.Permission) {
	dst.Id = p.Id
	dst.Code = p.Code
	dst.Name = p.Name
	dst.Description = p.Description
	dst.Category = p.Category
}

func (p *AuthPermissions) Convert(src dto.Permission) {
	p.Id = src.Id
	p.Code = src.Code
	p.Name = src.Name
	p.Description = src.Description
	p.Category = src.Category
}
