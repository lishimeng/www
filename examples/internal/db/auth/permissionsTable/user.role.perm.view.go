package permissionsTable

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/www/dto"
)

type AuthUserPermView struct {
	app.Pk
	RoleCode       string `orm:"column(role_code);null;readOnly"`
	SaasUserCode   string `orm:"column(saas_user_code);null"`
	PermissionCode string `orm:"column(permission_code);null"`
	IsGlobal       int    `orm:"column(is_global);null"`
}

func (u *AuthUserPermView) TableName() string {
	return "auth_user_perm_view"
}

func (u *AuthUserPermView) Transform(dst *dto.AuthUserPerm) {
	dst.Id = u.Id
	dst.RoleCode = u.RoleCode
	dst.SaasUserCode = u.SaasUserCode
	dst.PermissionCode = u.PermissionCode
	dst.IsGlobal = u.IsGlobal
}
