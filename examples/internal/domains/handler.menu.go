package domains

import (
	"errors"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/examples/internal/db/auth/permissionsTable"
)

func (h handler) GetPermRouters(saasUserCode string) (list []dto.Menu) {
	var models []permissionsTable.AuthUserRouterView
	_, _ = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthUserRouterView)).
		Filter("MenuGroup", www.SystemMenuGroup).
		Filter("Status", def.MenuStatusEnabled).
		Filter("SaasUserCode", saasUserCode).
		OrderBy("Sort", "CreateTime").
		All(&models)
	for _, model := range models {
		var one dto.Menu
		model.TransformMenuVO(&one)
		list = append(list, one)
	}
	return
}

func (h handler) GetCommonRouters() (list []dto.Menu) {
	var models []permissionsTable.AuthMenu
	_, _ = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthMenu)).
		Filter("MenuGroup", www.SystemMenuGroup).
		Filter("Status", def.MenuStatusEnabled).
		Filter("HasPerm", 0).
		OrderBy("Sort", "CreateTime").
		All(&models)
	for _, model := range models {
		var one dto.Menu
		model.TransformMenuVO(&one)
		list = append(list, one)
	}
	return
}

func (h handler) GetMenu(id int) (menu dto.Menu, err error) {
	var model permissionsTable.AuthMenu
	model.Id = id
	if err = app.GetOrm().Context.Read(&model); err != nil {
		return
	}
	model.TransformMenuVO(&menu)
	return
}

func (h handler) AddMenu(menu *dto.ReqMenuForm) (err error) {
	// 防重
	if len(menu.RouteName) != 0 {
		repeated := app.GetOrm().Context.QueryTable(new(permissionsTable.AuthMenu)).
			Filter("RouteName", menu.RouteName).
			Exist()
		if repeated {
			err = errors.New("路由名称重复")
			return
		}
	}
	var model permissionsTable.AuthMenu
	model.Convert(*menu)
	_, err = app.GetOrm().Context.Insert(&model)
	return
}

func (h handler) GetMenuOptionList(group def.MenuGroup) (list []dto.MenuOption, err error) {
	var models []permissionsTable.AuthMenu
	// 配置，不需要筛选权限
	_, err = app.GetOrm().Context.QueryTable(new(permissionsTable.AuthMenu)).
		Filter("MenuGroup", group). // todo: 允许菜单有隐藏子菜单
		Filter("Type", def.MenuTypeCatalog).
		All(&models)
	if err != nil {
		return
	}
	for _, model := range models {
		var one dto.MenuOption
		model.TransformOption(&one)
		list = append(list, one)
	}
	return
}

func (h handler) UpdateMenu(menu *dto.ReqMenuForm) (err error) {
	var model permissionsTable.AuthMenu
	model.Convert(*menu)
	_, err = app.GetOrm().Context.Update(&model)
	return
}

func _delMenu(tx persistence.TxContext, id int) (err error) {
	var children []permissionsTable.AuthMenu
	_, err = tx.Context.QueryTable(new(permissionsTable.AuthMenu)).
		Filter("ParentId", id).
		All(&children)
	if err != nil {
		return err
	}
	for _, child := range children {
		err = _delMenu(tx, child.Id)
		if err != nil {
			return err
		}
	}
	_, err = tx.Context.QueryTable(new(permissionsTable.AuthMenu)).
		Filter("Id", id).
		Delete()
	return nil
}

func (h handler) DeleteMenu(id int) (err error) {
	err = app.GetOrm().Transaction(func(tx persistence.TxContext) (e error) {
		// 递归删除
		e = _delMenu(tx, id)
		return
	})
	return
}

func (h handler) GetMenuList(filterKw string, filterProject def.MenuGroup) (list []dto.Menu, err error) {
	var models []permissionsTable.AuthMenu
	q := app.GetOrm().Context.QueryTable(new(permissionsTable.AuthMenu))
	if len(filterProject) > 0 {
		q = q.Filter("MenuGroup", filterProject)
	}
	if len(filterKw) > 0 {
		c := app.GetOrm().NewCondition()
		c = c.Or("Name__contains", filterKw).Or("RouteName__contains", filterKw)
		q = q.SetCond(c)
	}
	_, err = q.OrderBy("Sort", "CreateTime").All(&models)
	if err != nil {
		return
	}
	for _, model := range models {
		var one dto.Menu
		model.TransformMenuVO(&one)
		list = append(list, one)
	}
	return

}
