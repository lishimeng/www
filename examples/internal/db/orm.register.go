package db

import (
	"github.com/lishimeng/www/examples/internal/db/auth/permissionsTable"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
	"github.com/lishimeng/www/examples/internal/db/saasTable"
)

func RegisterTables() (models []interface{}) {
	models = append(models,
		new(permissionsTable.AuthMenu),
		new(permissionsTable.AuthPermissions),
		new(permissionsTable.AuthRolePermissions),
		new(permissionsTable.AuthRolesCategory),
		new(permissionsTable.AuthRoles),
		new(permissionsTable.AuthUserRoles),
		new(permissionsTable.AuthUserPermView),
		new(permissionsTable.AuthUserRouterView),
		new(usersTable.Auth2FA),
		new(usersTable.AuthIdentity),
		new(usersTable.AuthSecurityInfo),
		new(usersTable.AuthUserProfile),
		new(saasTable.SaasOrganization),
		new(saasTable.SaasUserInfo),
	)
	return
}
