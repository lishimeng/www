package www

import (
	"github.com/lishimeng/www/dto"
)

type DbRouteManager struct {
	domain DomainHandler
}

func NewDbRouteManager(db DomainHandler) RouteManager {
	var h RouteManager
	h = &DbRouteManager{
		domain: db,
	}
	return h
}

func (m *DbRouteManager) GetUserRouters(saasUserCode string) (routers []dto.Menu) {
	permissionMap := make(map[int]byte)

	commonRouters := m.domain.GetCommonRouters()
	for _, router := range commonRouters {
		if _, ok := permissionMap[router.GetId()]; !ok {
			permissionMap[router.GetId()] = 0x01
			routers = append(routers, router)
		}
	}
	if len(saasUserCode) == 0 {
		return
	}

	userRouters := m.domain.GetPermRouters(saasUserCode)
	for _, router := range userRouters {
		if _, ok := permissionMap[router.GetId()]; !ok {
			permissionMap[router.GetId()] = 0x01
			routers = append(routers, router)
		}
	}
	return
}
