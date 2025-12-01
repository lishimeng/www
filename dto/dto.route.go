package dto

type MenuRouteMeta struct {
	Title      string            `json:"title,omitempty"`
	Icon       string            `json:"icon,omitempty"`
	Hidden     bool              `json:"hidden,omitempty"`
	AlwaysShow bool              `json:"alwaysShow,omitempty"`
	KeepAlive  bool              `json:"keepAlive,omitempty"`
	Params     map[string]string `json:"params,omitempty"` // 查询参数
}

type MenuRoute struct {
	Children  []*MenuRoute   `json:"children,omitempty"`
	Component string         `json:"component,omitempty"`
	Meta      *MenuRouteMeta `json:"meta,omitempty"`
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Redirect  string         `json:"redirect,omitempty"`
}
