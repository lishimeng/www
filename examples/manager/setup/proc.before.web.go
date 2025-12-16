package setup

import (
	"context"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/examples/internal/domains"
)

func ComponentBeforeWeb(_ context.Context) error {
	// manager 实施配置，管理账号、身份以及saas组织
	www.SystemScope = def.Manager
	www.SystemMenuGroup = "m_manager"
	domains.InitManagers()
	def.QueryPageNum = "pageNum"
	def.QueryPageSize = "pageSize"
	def.DefaultQueryPageSize = 10
	return nil
}
