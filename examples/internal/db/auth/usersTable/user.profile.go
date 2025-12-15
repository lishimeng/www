package usersTable

import (
	"github.com/lishimeng/app-starter"
)

type AuthUserProfile struct {
	app.Pk
	Code   string `orm:"column(code)"` // security code 安全码
	Name   string `orm:"column(name)"`
	Avatar string `orm:"column(avatar)"`

	app.TableChangeInfo
}
