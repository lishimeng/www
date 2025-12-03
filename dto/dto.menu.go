package dto

import (
	"fmt"

	"github.com/lishimeng/www/def"
	"github.com/tdewolff/parse/v2/strconv"
)

type MenuOption struct {
	Value    string        `json:"value"` // 使用string类型以匹配前端OptionType
	Label    string        `json:"label"`
	ParentId string        `json:"-"` // 内部使用string以便与Value类型一致
	Children []*MenuOption `json:"children"`
}

func (m *MenuOption) AddChildren(m2 *MenuOption) {
	m.Children = append(m.Children, m2)
}

func (m *MenuOption) GetId() string {
	return m.Value
}

func (m *MenuOption) GetParentId() string {
	return m.ParentId
}

type MenuParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ReqMenuForm struct {
	Id        string        `json:"id"`
	ParentId  string        `json:"parentId"`
	Name      string        `json:"name"`
	Visible   int           `json:"visible"`
	Icon      string        `json:"icon"`
	Sort      int           `json:"sort"`
	RouteName string        `json:"routeName"`
	RoutePath string        `json:"routePath"`
	Component string        `json:"component"`
	Type      int           `json:"type"`
	Perm      string        `json:"perm"`
	HasPerm   bool          `json:"hasPerm"`
	KeepAlive int           `json:"keepAlive"`
	Params    []MenuParam   `json:"params"`
	Group     def.MenuGroup `json:"group"` // 菜单所属程序组 - project(app) - menu category
}

type Menu struct {
	Id         string        `json:"id"`
	ParentId   string        `json:"parentId"`
	Name       string        `json:"name"`
	Visible    int           `json:"visible"`
	Icon       string        `json:"icon"`
	Sort       int           `json:"sort"`
	RouteName  string        `json:"routeName"`
	RoutePath  string        `json:"routePath"`
	Component  string        `json:"component"`
	Type       int           `json:"type"`
	Perm       string        `json:"perm"`
	HasPerm    bool          `json:"hasPerm"`
	KeepAlive  int           `json:"keepAlive"`
	Params     []MenuParam   `json:"params"`
	Hidden     int           `json:"hidden"`
	Group      def.MenuGroup `json:"group"` // 菜单所属程序组
	Children   []*Menu       `json:"children"`
	CreateTime string        `json:"createTime"`
	UpdateTime string        `json:"updateTime"`
}

func ConvertInt(s string) (i int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			i = 0
		}
	}()
	i64, _ := strconv.ParseInt([]byte(s))
	i = int(i64)
	return i
}

func (m *Menu) AddChildren(m2 *Menu) {
	m.Children = append(m.Children, m2)
}

func (m *Menu) GetId() int {
	return ConvertInt(m.Id)
}

func (m *Menu) GetParentId() int {
	return ConvertInt(m.ParentId)
}
