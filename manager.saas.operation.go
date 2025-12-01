package www

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/www/dto"
)

type DbSaasOperationManager struct {
	domain SaasOperationDomainHandler
}

func NewDbSaasOperationManager(domain SaasOperationDomainHandler) SaasOperationManager {
	return &DbSaasOperationManager{domain: domain}
}

func (m *DbSaasOperationManager) CreateSaas(userCode string, orgCode string) (one dto.SaasUser, err error) {
	one, err = m.domain.CreateSaasUser(userCode, orgCode)
	return
}

func (m *DbSaasOperationManager) SaasList(orgCode string, pager *app.Pager[dto.SaasUser]) (err error) {
	return
}
