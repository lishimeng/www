package usersTable

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/www/def"
)

type Auth2FA struct { // 双因子验证
	app.Pk
	Code             string          `orm:"column(code)"`               // 编号
	FaType           def.Auth2FaType `orm:"column(fa_type)"`            // 2FA类型
	FaDeviceSn       string          `orm:"column(fa_device_sn)"`       // 2FA设备编号(手机号/OPT验证器)
	SecurityUserCode string          `orm:"column(security_user_code)"` // 用户
	app.TableChangeInfo
}

func (auth *Auth2FA) TableName() string {
	return "auth_2fa"
}
