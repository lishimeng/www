package www

type DbAuthManager struct {
	domain AuthDomainHandler
}

func NewDbAuthManager(db AuthDomainHandler) AuthManager {
	var h AuthManager
	h = &DbAuthManager{domain: db}
	return h
}
