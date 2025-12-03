package saasTable

import "github.com/lishimeng/app-starter"

type SaasOrganizationUser struct {
	app.Pk
	OrganizationCode string `orm:"column(organization_code)"`  // 组织编号
	SecurityUserCode string `orm:"column(security_user_code)"` // 用户编号
	IsAdmin          int    `orm:"column(is_admin)"`           // 是组织内的管理员(1:是, 0:否)
	app.TableInfo           // 只有添加和删除(绑定/解绑)
}
