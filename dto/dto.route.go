package dto

type MenuRouteMeta struct {
	Title      string            `json:"title"`
	Icon       string            `json:"icon"`
	Hidden     bool              `json:"hidden"`
	AlwaysShow bool              `json:"alwaysShow"`
	KeepAlive  bool              `json:"keepAlive"`
	Params     map[string]string `json:"params"` // 查询参数
}

type MenuRoute struct {
	Children  []*MenuRoute   `json:"children,omitempty"`
	Component string         `json:"component"`
	Meta      *MenuRouteMeta `json:"meta"`
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Redirect  string         `json:"redirect,omitempty"`
}
