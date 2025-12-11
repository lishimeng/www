package www

import (
	"errors"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www/dto"
)

type DbSaasOperationManager struct {
	domain SaasOperationDomainHandler
}

func NewDbSaasOperationManager(domain SaasOperationDomainHandler) SaasOperationManager {
	return &DbSaasOperationManager{domain: domain}
}

func saasUserCode(orgCode, userCode string) string {
	return orgCode + userCode
}

func (m *DbSaasOperationManager) CreateSaas(userCode string, orgCode string, name string) (one dto.SaasUser, err error) {
	err = app.Transaction(func(tx persistence.TxContext) (e error) {
		_, e = m.domain.TxCreateSaasOrg(tx, orgCode, name, userCode)
		if e != nil {
			log.Info("创建组织失败: %s", orgCode)
			return
		}
		one = dto.SaasUser{
			Code:           saasUserCode(orgCode, userCode),
			Name:           "Admin",
			EnterpriseCode: orgCode,
			SecurityCode:   userCode,
		}
		e = m.domain.TxAddOrgUser(tx, &one)
		return
	})
	return
}

func (m *DbSaasOperationManager) SaasList(orgCode string, pager *app.Pager[dto.SaasUser]) (err error) {
	err = m.domain.SaasUserPage(orgCode, pager)
	return
}

func (m *DbSaasOperationManager) AllOrgs() (list []dto.SaasOrganization, err error) {
	return m.domain.GetAllOrgs()
}

func (m *DbSaasOperationManager) JoinOrganization(orgCode string, userCode string, userName string) (u dto.SaasUser, err error) {
	err = app.Transaction(func(tx persistence.TxContext) (e error) {
		_, e = m.domain.TxGetOrg(tx, orgCode)
		if e != nil {
			log.Info("组织不存在: %s", orgCode)
			e = errors.New("组织不存在: " + orgCode)
			return
		}
		_, e = m.domain.TxGetSaasUser(tx, saasUserCode(orgCode, userCode))
		if e == nil {
			log.Info("用户已加入组织: %s", userCode)
			e = errors.New("用户已加入组织: " + userCode)
			return
		}
		u = dto.SaasUser{
			Code:           saasUserCode(orgCode, userCode),
			Name:           userName,
			EnterpriseCode: orgCode,
			SecurityCode:   userCode,
		}
		e = m.domain.TxAddOrgUser(tx, &u)
		if e != nil {
			log.Info("加入组织失败: %s", orgCode)
			return
		}
		return
	})
	return
}

func (m *DbSaasOperationManager) LeaveOrganization(orgCode string, userCode string) (err error) {
	err = app.Transaction(func(tx persistence.TxContext) (e error) {
		var org dto.SaasOrganization
		org, e = m.domain.TxGetOrg(tx, orgCode)
		if e != nil {
			log.Info("组织不存在: %s", orgCode)
			e = errors.New("组织不存在: " + orgCode)
			return
		}
		if org.Admin == userCode {
			log.Info("组织管理员不能退出组织: %s", orgCode)
			e = errors.New("组织管理员不能退出组织: " + orgCode)
			return
		}
		_, e = m.domain.TxGetSaasUser(tx, saasUserCode(orgCode, userCode))
		if e != nil {
			log.Info("企业用户不存在: %s", userCode)
			e = errors.New("企业用户不存在: " + userCode)
			return
		}
		e = m.domain.TxDelOrgUser(tx, saasUserCode(orgCode, userCode))
		return
	})
	return
}
