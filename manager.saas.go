package www

import "github.com/lishimeng/www/dto"

type DbSaasManager struct {
	domain SaasDomainHandler
}

func NewDbSaasManager(domain SaasDomainHandler) SaasManager {
	return &DbSaasManager{domain: domain}
}

func (m *DbSaasManager) UserSaasOrgs(userCode string) (list []dto.SaasUser, err error) {
	return m.domain.GetUserSaasOrgs(userCode)
}

func (m *DbSaasManager) SaasOrg(code string) (one dto.SaasOrganization, err error) {
	return m.domain.GetSaasOrg(code)
}
