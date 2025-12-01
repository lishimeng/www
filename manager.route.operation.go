package www

import (
	"errors"
	"strings"

	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
)

type DbRouteOperationManager struct {
	domain DomainHandler
}

func NewDbRouteOperationManager(db DomainHandler) RouteOperationManager {
	var h RouteOperationManager
	h = &DbRouteOperationManager{
		domain: db,
	}
	return h
}

func (mom *DbRouteOperationManager) GetRouteOptions(group string) (list []dto.MenuOption, err error) {
	menuGroup := SystemMenuGroup
	if len(group) > 0 {
		menuGroup = def.MenuGroup(group)
	}
	list, err = mom.domain.GetMenuOptionList(menuGroup)
	return
}

func (mom *DbRouteOperationManager) GetRoute(id int) (menu dto.Menu, err error) {
	return mom.domain.GetMenu(id)
}

func (mom *DbRouteOperationManager) GetProjectRoutes(kw, projectCode string) (list []dto.Menu) {
	// 获取menuGroup参数，如果未指定则使用当前程序的MenuCategory
	menuGroup := SystemMenuGroup
	if len(projectCode) > 0 {
		menuGroup = def.MenuGroup(projectCode)
	}
	list, _ = mom.domain.GetMenuList(kw, menuGroup)
	return
}

func (mom *DbRouteOperationManager) DelRoute(id int) (err error) {
	err = mom.domain.DeleteMenu(id)
	return
}

func (mom *DbRouteOperationManager) EditRoute(dto *dto.ReqMenuForm) (err error) {
	if err := validateMenu(dto); err != nil {
		return err
	}

	// 如果MenuGroup未设置，自动从SysInformation获取
	// 注意：在admin中管理其他程序的菜单时，MenuGroup应该已经通过前端指定
	if len(dto.Group) == 0 {
		dto.Group = SystemMenuGroup
	}
	err = mom.domain.UpdateMenu(dto)
	return
}

func (mom *DbRouteOperationManager) AddRoute(dto *dto.ReqMenuForm) (err error) {
	err = validateMenu(dto)
	if err != nil {
		return
	}
	// 如果MenuGroup未设置，自动从SysInformation获取
	// 注意：在admin中管理其他程序的菜单时，MenuGroup应该已经通过前端指定
	if len(dto.Group) == 0 {
		dto.Group = SystemMenuGroup
	}
	err = mom.domain.AddMenu(dto)
	return
}

var errRootRoutePath = errors.New("顶层菜单的路由路径必须以'/'开头")

// validateMenu 验证菜单数据
func validateMenu(menu *dto.ReqMenuForm) error {
	// 如果是顶层菜单（parentId = 0），路由路径必须以"/"开头
	if menu.ParentId == "0" && menu.RoutePath != "" {
		if !strings.HasPrefix(menu.RoutePath, "/") {
			return errRootRoutePath
		}
	}
	return nil
}
