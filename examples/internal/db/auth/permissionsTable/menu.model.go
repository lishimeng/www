package permissionsTable

import (
	"encoding/json"
	"fmt"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
	"strings"
	"time"

	"github.com/lishimeng/app-starter"
)

type AuthMenu struct {
	app.Pk
	Name      string        `orm:"column(name);size(100)"`        // 菜单名称（显示）
	RoutePath string        `orm:"column(route_path);size(200)"`  // 路由路径
	Component string        `orm:"column(component);size(200)"`   // 前端组件路径
	Type      int           `orm:"column(type);size(100)"`        // 类型：目录/页面/按钮（组件）/api
	Icon      string        `orm:"column(icon);size(100)"`        // 图标
	ParentId  int           `orm:"column(parent_id);default(0)"`  // 父级菜单ID
	RouteName string        `orm:"column(route_name);size(100)"`  // 路由名称（唯一）
	KeepAlive int           `orm:"column(keep_alive);default(0)"` // 是否缓存
	Sort      int           `orm:"column(sort);default(0)"`       // 排序
	Visible   int           `orm:"column(visible);default(1)"`    // 是否在菜单中显示
	Params    string        `orm:"column(params);size(200)"`      // 查询参数
	Perm      string        `orm:"column(perm);size(200)"`        // 权限
	HasPerm   int           `orm:"column(has_perm)"`              // 1:控制权限, 0:公共menu
	MenuGroup def.MenuGroup `orm:"column(menu_group);null"`       // 菜单所属程序组, 在project中标识 没有公共菜单
	app.TableChangeInfo
}

func (am *AuthMenu) TransformMenuVO(dst *dto.Menu) {
	dst.Id = fmt.Sprintf("%d", am.Id)
	dst.ParentId = fmt.Sprintf("%d", am.ParentId)
	dst.Name = am.Name
	dst.Type = am.Type
	dst.RoutePath = "/" + am.RoutePath
	if len(am.RouteName) == 0 {
		dst.RouteName = dst.RoutePath
	} else {
		dst.RouteName = am.RouteName
	}
	dst.Component = am.Component
	dst.Perm = am.Perm
	dst.HasPerm = am.HasPerm == 1
	dst.Visible = am.Visible
	dst.Sort = am.Sort
	dst.Icon = am.Icon
	dst.KeepAlive = am.KeepAlive
	dst.Group = am.MenuGroup
	dst.CreateTime = am.CreateTime.UTC().Format(time.RFC3339)
	dst.UpdateTime = am.UpdateTime.UTC().Format(time.RFC3339)
	dst.Status = am.Status

	if len(am.Params) > 0 {
		bs := []byte(am.Params)
		_ = json.Unmarshal(bs, &dst.Params)
	}
}

func (am *AuthMenu) Convert(src dto.ReqMenuForm) {
	// 如果Id不为空且不为"0"，设置主键Id（用于更新）
	// 创建新菜单时，Id为空，数据库会自动分配自增ID
	if src.Id != "" && src.Id != "0" {
		am.Id = dto.ConvertInt(src.Id)
	}
	// 如果Id为"0"，表示创建新菜单，不设置Id让数据库自增
	am.ParentId = dto.ConvertInt(src.ParentId)
	am.Name = src.Name
	am.Type = src.Type
	am.RouteName = src.RouteName
	am.RoutePath = strings.TrimPrefix(src.RoutePath, "/")
	am.Component = src.Component
	am.Visible = src.Visible
	am.Sort = src.Sort
	am.Icon = src.Icon
	am.KeepAlive = src.KeepAlive
	if len(src.Params) > 0 {
		bs, _ := json.Marshal(src.Params)
		am.Params = string(bs)
	}

	// 处理权限相关字段
	if src.HasPerm {
		// 启用权限：设置HasPerm=1，保留Perm字段
		am.HasPerm = 1
		am.Perm = src.Perm
	} else {
		// 关闭权限：清空Perm字段，设置HasPerm=0
		am.HasPerm = 0
		am.Perm = ""
	}
	// 设置Group字段（如果前端传了则使用，否则不设置，由创建/更新逻辑自动设置）
	if len(src.Group) > 0 {
		am.MenuGroup = src.Group
	}
	am.Status = src.Status
}

func (am *AuthMenu) TransformOption(dst *dto.MenuOption) {
	dst.Value = fmt.Sprintf("%d", am.Id)
	if am.ParentId == 0 {
		dst.ParentId = ""
	} else {
		dst.ParentId = fmt.Sprintf("%d", am.ParentId)
	}
	dst.Label = am.Name
}
