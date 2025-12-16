package permissionsTable

type AuthUserRouterView struct {
	AuthMenu
	SaasUserCode string `orm:"column(saas_user_code);null;readOnly"`
	RoleCode     string `orm:"column(role_code);null;readOnly"`
}

func (u *AuthUserRouterView) Transform(dto *AuthUserRouterView) {
	return
}
