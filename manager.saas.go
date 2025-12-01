package www

type DbSaasManager struct {
	domain SaasDomainHandler
}

func NewDbSaasManager(domain SaasDomainHandler) SaasManager {
	return &DbSaasManager{domain: domain}
}
