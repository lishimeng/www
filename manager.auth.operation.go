package www

import (
	"github.com/lishimeng/www/dto"
)

type DbAuthOperationManager struct {
	domain AuthOperationDomainHandler
}

func NewDbAuthOperationManager(domain AuthOperationDomainHandler) AuthOperationManager {
	var h AuthOperationManager
	h = &DbAuthOperationManager{domain: domain}
	return h
}

func (om *DbAuthOperationManager) SecurityUser(code string) (one dto.SecurityInfo, err error) {

	one, err = om.domain.GetSecurityInfo(code)
	return
}
