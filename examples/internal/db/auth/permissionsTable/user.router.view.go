package permissionsTable

type AuthUserRouterView struct {
	AuthMenu
	UserCode string `orm:"column(user_code);null;readOnly"`
	RoleCode string `orm:"column(role_code);null;readOnly"`
}

func (u *AuthUserRouterView) Transform(dto *AuthUserRouterView) {
	return
}
