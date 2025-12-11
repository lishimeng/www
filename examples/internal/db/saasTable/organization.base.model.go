package saasTable

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/www/dto"
)

// SaasOrganization saas组织主表, 只提供编号与admin_user的绑定, 附属信息表在业务表中按需实现, 建议saas相关业务表中冗余记录org_code
type SaasOrganization struct {
	app.Pk
	Code      string `orm:"column(code);unique"`
	Name      string `orm:"column(name);unique"`
	AdminUser string `orm:"column(admin_user);null"` // 管理员用户
	app.TableChangeInfo
}

func (so SaasOrganization) Transform(dst *dto.SaasOrganization) {
	dst.Id = so.Id
	dst.Code = so.Code
	dst.Name = so.Name
	dst.Admin = so.AdminUser
}
