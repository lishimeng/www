package def

type MenuGroup string // menu所属程序

const (
	MenuHide = iota
	MenuShow
)

const (
	MenuStatusDisabled = iota
	MenuStatusEnabled
)

const (
	MenuTypeApi = iota
	MenuTypeMenu
	MenuTypeCatalog
)
