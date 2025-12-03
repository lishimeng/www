package usersTable

import (
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
	"time"

	"github.com/lishimeng/app-starter"
)

// AuthIdentity 登录标识
type AuthIdentity struct {
	app.Pk
	Code         string           `orm:"column(code);unique"`   // 登录标识
	IdentityType def.IdentityType `orm:"column(identity_type)"` // 类型
	SecurityCode string           `orm:"column(security_code)"` // 账号，用于登录
	app.TableChangeInfo
}

func (ai *AuthIdentity) Transform(dst *dto.Identity) {
	dst.Id = ai.Id
	dst.Code = ai.Code
	dst.IdentityType = int(ai.IdentityType)
	dst.SecurityCode = ai.SecurityCode
	dst.Status = ai.Status
	dst.CreateTime = ai.CreateTime.UTC().Format(time.RFC3339)
	dst.UpdateTime = ai.UpdateTime.UTC().Format(time.RFC3339)
}
