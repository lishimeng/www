package menu

import (
	"strings"
	"unicode"

	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
)

func ToCamelCase(s string) string {
	s = strings.TrimPrefix(s, "/")
	if s == "" {
		return ""
	}
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-'
	})
	for i, word := range words {
		if word == "" {
			continue
		}
		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}
	return strings.Join(words, "")
}

func buildParam(ls []dto.MenuParam) map[string]string {
	if len(ls) == 0 {
		return nil
	}
	m := make(map[string]string)
	for _, v := range ls {
		m[v.Key] = v.Value
	}
	return m
}

func buildRedirect(path string, child *dto.MenuRoute) string {
	if len(child.Redirect) != 0 {
		return child.Redirect
	} else {
		return path + "/" + child.Path
	}
}

func BuildRoutes(tree []*dto.Menu, path string) (b []*dto.MenuRoute) {
	// 初始化为空切片，确保不会返回nil
	b = []*dto.MenuRoute{}
	// 形成树才能进行下一步构建
	// 处理一级目录下直接的菜单
	// 处理redirect
	isPrimary := path == ""
	for _, n := range tree {
		if isPrimary {
			if n.Type == def.MenuTypeCatalog {
				children := BuildRoutes(n.Children, n.RoutePath)
				if len(children) == 0 {
					continue
				}
				b = append(b, &dto.MenuRoute{
					Name:      n.RoutePath,
					Path:      n.RoutePath,
					Component: "Layout",
					Redirect:  buildRedirect(path, children[0]),
					Children:  children,
					Meta: &dto.MenuRouteMeta{
						Title:      n.Name,
						Icon:       n.Icon,
						Hidden:     n.Visible == def.MenuHide,
						AlwaysShow: true,
						Params:     buildParam(n.Params),
					},
				})
			} else {
				child := &dto.MenuRoute{
					Name:      ToCamelCase(n.RoutePath),
					Path:      "",
					Component: n.Component,
					Meta: &dto.MenuRouteMeta{
						Title:      n.Name,
						Icon:       n.Icon,
						Hidden:     n.Visible == def.MenuHide,
						AlwaysShow: false,
						Params:     buildParam(n.Params),
					},
					Children: nil,
				}
				b = append(b, &dto.MenuRoute{
					Name:      n.RoutePath,
					Path:      n.RoutePath,
					Redirect:  "",
					Component: "Layout",
					Meta:      nil,
					Children:  []*dto.MenuRoute{child},
				})
			}
		} else {
			// 多级嵌套目录
			if n.Type == def.MenuTypeCatalog {
				children := BuildRoutes(n.Children, path+n.RoutePath)
				if len(children) == 0 {
					continue
				}
				b = append(b, &dto.MenuRoute{
					Name:      ToCamelCase(n.RoutePath),
					Path:      strings.TrimPrefix(n.RoutePath, "/"),
					Component: "",
					Redirect:  buildRedirect(path+n.RoutePath, children[0]),
					Children:  children,
					Meta: &dto.MenuRouteMeta{
						Title:      n.Name,
						Icon:       n.Icon,
						Hidden:     n.Visible == def.MenuHide,
						AlwaysShow: true,
						Params:     buildParam(n.Params),
					}})
			} else {
				b = append(b, &dto.MenuRoute{
					Name:      ToCamelCase(n.RoutePath),
					Path:      strings.TrimPrefix(n.RoutePath, "/"),
					Component: n.Component,
					Meta: &dto.MenuRouteMeta{
						Title:      n.Name,
						Icon:       n.Icon,
						Hidden:     n.Visible == def.MenuHide,
						AlwaysShow: false,
						Params:     buildParam(n.Params),
					}})
			}
		}
	}
	return
}
