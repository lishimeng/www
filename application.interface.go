package www

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
)

type RoleManager interface {
	UserRoles(userCode string) (role []dto.UserRole, err error)
}
type RoleDomainHandler interface {
	GetUserRoleList(userCode string) (role []dto.UserRole, err error)
}

type RoleOperationManager interface {
	All() (list []dto.Role, err error)
	AddRole(role *dto.Role) (err error)
	EditRole(role *dto.Role) (err error)
	DelRole(roleCode string) (err error)
	One(roleCode string) (one dto.Role, err error)
	BindUser(userCode, roleCode string) (ur dto.UserRole, err error)
	UnbindUser(userCode, roleCode string) (err error)
	RoleUsers(roleCode string) (list []dto.UserRole, err error)
}
type RoleOperationDomainHandler interface {
	GetAllRole() (list []dto.Role, err error)
	CreateRole(role *dto.Role) (err error)
	UpdateRole(role *dto.Role) (err error)
	DelRole(roleCode string) (err error)
	GetRoleByCode(roleCode string) (role dto.Role, err error)
	CreateUserRole(userCode, roleCode string) (ur dto.UserRole, err error)
	DeleteUserRole(userCode, roleCode string) (err error)
	GetUserRole(filterUserCode, filterRoleCode string) (list []dto.UserRole, err error)
}

type PermManager interface {
	GetUserPerms(user string) (perms []dto.AuthUserPerm) // 用户的权限列表
}
type PermDomainHandler interface {
	GetPermissionList(filterUserCode string, filterProjectCode string) (list []dto.AuthUserPerm, err error)
}

type PermOperationManager interface {
	CreatePerm(dto *dto.Permission) (err error)
	One(code string) (dto dto.Permission, err error)
	DelPerm(code string) (err error)
	EditPerm(code string, dto *dto.Permission) (err error)
	AllPerms() (perms []dto.Permission)
	BindRole(roleCode, permCome string) (err error)
	UnbindRole(roleCode, permCome string) (err error)
	GetRolePerms(roleCode string) (perms []dto.Permission)
}
type PermOperationHandler interface {
	TxCreatePerm(tx persistence.TxContext, dto *dto.Permission) (err error)
	TxUpdatePerm(tx persistence.TxContext, dto *dto.Permission) (err error)
	TxDeletePerm(tx persistence.TxContext, code string) (err error)
	GetAllPerms() (perms []dto.Permission)
	GetPermsByRole(roleCode string) (perms []dto.Permission)
	GetPermByCode(code string) (dto dto.Permission, err error)
	TxCreateRolePermNoChk(tx persistence.TxContext, dto *dto.AuthRolePerm) (err error)
	TxDeleteRolePermNoChk(tx persistence.TxContext, dto *dto.AuthRolePerm) (err error)
}

type IdentityDomainHandler interface {
	TxGetIdentity(tx persistence.TxContext, identityCode string) (one dto.Identity, err error)
	TxGetSecurity(tx persistence.TxContext, securityUserCode string) (one dto.SecurityInfo, err error)
	TxCreateIdentity(tx persistence.TxContext, securityUserCode, identityCode string, identityType def.IdentityType) (one dto.Identity, err error)
	TxDelIdentity(tx persistence.TxContext, code string) (err error)
	TxGetIdentities(ctx persistence.TxContext, securityUserFilter string, limits ...int) ([]dto.Identity, error)
	GetIdentityPage(pager *app.Pager[dto.Identity]) (err error)
	DetectIdentityType(identityCode string) def.IdentityType
}

type MenuDomainHandler interface {
	GetUserRouters(userCode string) (list []dto.Menu)
	GetCommonRouters() (list []dto.Menu)
	GetMenu(id int) (menu dto.Menu, err error)
	AddMenu(menu *dto.ReqMenuForm) (err error)
	GetMenuOptionList(group def.MenuGroup) (list []dto.MenuOption, err error)
	UpdateMenu(menu *dto.ReqMenuForm) (err error)
	DeleteMenu(id int) (err error)
	GetMenuList(filterKw string, filterProject def.MenuGroup) (list []dto.Menu, err error)
}

type RouteManager interface {
	GetUserRouters(userCode string) (list []dto.Menu)
}

type RouteOperationManager interface {
	AddRoute(*dto.ReqMenuForm) (err error)
	EditRoute(*dto.ReqMenuForm) (err error)
	DelRoute(id int) (err error)
	GetRoute(id int) (menu dto.Menu, err error)
	GetRouteOptions(menuGroup string) (list []dto.MenuOption, err error)
	GetProjectRoutes(kw, projectCode string) (list []dto.Menu)
}

type AuthManager interface {
}
type AuthDomainHandler interface {
}

type AuthOperationManager interface {
	SecurityUser(code string) (dto.SecurityInfo, error)
}
type AuthOperationDomainHandler interface {
	GetSecurityInfo(code string) (dto.SecurityInfo, error)
}

type IdentityManager interface {
	Bind(identityCode, securityUserCode string, identityType def.IdentityType) (err error)
	Unbind(identityCode string, securityUserCode string) (err error)
	Page(pager *app.Pager[dto.Identity]) (err error)
	UserIdentities(userCode string) (list []dto.Identity, err error)
	DetectIdentityType(identityCode string) def.IdentityType
}

type SaasManager interface {
	UserSaasOrgs(userCode string) (list []dto.SaasUser, err error)
	SaasOrg(code string) (one dto.SaasOrganization, err error)
}
type SaasDomainHandler interface {
	GetUserSaasOrgs(userCode string) (list []dto.SaasUser, err error)
	GetSaasOrg(code string) (one dto.SaasOrganization, err error)
}

type SaasOperationManager interface {
	CreateSaas(userCode string, orgCode string, name string) (u dto.SaasUser, err error)
	SaasList(orgCode string, pager *app.Pager[dto.SaasUser]) (err error)
	AllOrgs() (list []dto.SaasOrganization, err error)
	JoinOrganization(orgCode string, userCode string, userName string) (u dto.SaasUser, err error)
	LeaveOrganization(orgCode string, userCode string) (err error)
}
type SaasOperationDomainHandler interface {
	TxCreateSaasOrg(tx persistence.TxContext, orgCode string, name string, admin string) (org dto.SaasOrganization, err error)
	TxAddOrgUser(tx persistence.TxContext, dto *dto.SaasUser) (err error)
	TxGetOrg(tx persistence.TxContext, orgCode string) (org dto.SaasOrganization, err error)
	TxDelOrgUser(tx persistence.TxContext, code string) (err error)
	GetAllOrgs() (list []dto.SaasOrganization, err error)
	SaasUserPage(orgCode string, pager *app.Pager[dto.SaasUser]) (err error)
	TxGetSaasUser(tx persistence.TxContext, code string) (u dto.SaasUser, err error)
}

type DomainHandler interface {
	IdentityDomainHandler
	MenuDomainHandler
	PermDomainHandler
	PermOperationHandler
	RoleDomainHandler
	RoleOperationDomainHandler
	AuthDomainHandler
	AuthOperationDomainHandler
	SaasDomainHandler
	SaasOperationDomainHandler
}
