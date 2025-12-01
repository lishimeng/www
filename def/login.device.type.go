package def

type LoginClient string

const (
	PC      LoginClient = "pc"
	WxMini  LoginClient = "wxMini"
	AppIos  LoginClient = "ios"
	Android LoginClient = "android"
	Harmony LoginClient = "harmony"
)
