package www

import "github.com/lishimeng/www/dto"

type DbPermManager struct {
	domain DomainHandler
}

func NewPermManager(db DomainHandler) PermManager {
	var h PermManager
	h = &DbPermManager{domain: db}
	return h
}

func (pm *DbPermManager) GetUserPerms(user string) (perms []dto.AuthUserPerm) {
	list, err := pm.domain.GetPermissionList(user, "")
	if err == nil {
		perms = list
	}
	return
}
