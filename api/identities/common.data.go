package identities

import (
	"fmt"

	"github.com/lishimeng/www/def"
)

type IdentityForm struct {
	IdentityCode string `json:"identityCode,omitempty"`
	UserCode     string `json:"userCode,omitempty"`
	IdentityType *int   `json:"identityType,omitempty"` // 身份类型，可选，如果不提供则自动判断
}

// GetIdentityTypeName 获取身份类型名称
func GetIdentityTypeName(identityType int) string {
	switch def.IdentityType(identityType) {
	case def.LoginCode:
		return "登录码"
	case def.LoginPhone:
		return "手机号"
	case def.LoginEmail:
		return "邮箱"
	case def.LoginWechatUnionId:
		return "微信UnionId"
	case def.LoginWechatOpenId:
		return "微信OpenId"
	default:
		return fmt.Sprintf("未知类型(%d)", identityType)
	}
}

// GetPlatformName 获取平台名称
func GetPlatformName(platform int) string {
	switch def.UserPlatform(platform) {
	case def.AppUser:
		return "App用户"
	case def.SaasUser:
		return "Saas用户"
	case def.AdminUser:
		return "管理员"
	default:
		return fmt.Sprintf("未知平台(%d)", platform)
	}
}
