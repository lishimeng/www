package usersTable

import (
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/passwords"
	"time"

	"github.com/lishimeng/app-starter"
)

type AuthSecurityInfo struct {
	app.Pk
	Code           string           `orm:"column(code);unique"`      // userCode; SecurityCode
	Password       string           `orm:"column(password)"`         // 密码是底线,与其他类型授权区别开对待
	Platform       def.UserPlatform `orm:"column(platform)"`         // 指定platform后禁止跨平台串账号
	SecondFaEnable int              `orm:"column(second_fa_enable)"` // 启动2FA
	SecondFaCode   string           `orm:"column(second_fa_code)"`   // 2FA设备编号
	app.TableChangeInfo
}

func (si *AuthSecurityInfo) Transform(dst *dto.SecurityInfo) {
	dst.Id = si.Id
	dst.Code = si.Code
	dst.Platform = int(si.Platform)
	dst.SecondFaEnable = si.SecondFaEnable
	dst.SecondFaCode = si.SecondFaCode
	dst.Status = si.Status
	dst.CreateTime = si.CreateTime.UTC().Format(time.RFC3339)
	dst.UpdateTime = si.UpdateTime.UTC().Format(time.RFC3339)
}

func (si *AuthSecurityInfo) getPwdTs() int64 {
	return si.CreateTime.Unix()
}

func (si *AuthSecurityInfo) getSalt() string {
	return si.Code
}

func (si *AuthSecurityInfo) EncodingPassword(plaintext string) string {
	return passwords.Encode(plaintext, si.getPwdTs(), si.getSalt())
}

func (si *AuthSecurityInfo) ValidatePassword(plaintext string) bool {
	return passwords.Validate(plaintext, si.getPwdTs(), si.getSalt(), si.Password)
}
