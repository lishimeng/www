package security

import (
	"errors"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
	"github.com/lishimeng/www/examples/internal/domains"
)

func UpdatePsw(securityCode string, newPassword string, oldPswValidator func(security usersTable.AuthSecurityInfo) bool) (err error) {
	err = app.GetOrm().Transaction(func(tx persistence.TxContext) (e error) {
		security, e := domains.TxGetSecurity(tx, securityCode)
		if e != nil {
			log.Info("get security info failed")
			log.Info(e)
			return
		}
		if oldPswValidator == nil {
			log.Info("old password validator nil")
			err = errors.New("old password validator nil")
			return
		}
		if !oldPswValidator(security) {
			log.Info("old password is invalid")
			err = errors.New("旧密码不正确")
			return
		}
		security.Password = security.EncodingPassword(newPassword)
		e = domains.TxUpdateSecurity(tx, security, "Password")
		return
	})
	if err != nil {
		log.Info("update security info failed")
		log.Info(err)
		err = errors.New("更新密码失败")
	}
	return
}

func CreateSecurityUser(identity string,
	identityType def.IdentityType,
	platform def.UserPlatform) (identityInfo dto.Identity, user dto.SecurityInfo, err error) {
	err = app.GetOrm().Transaction(func(tx persistence.TxContext) (e error) {
		identityInfo, e = domains.TxGetIdentity(tx, identity)
		if e == nil {
			e = errors.New(fmt.Sprintf("identity[%s]已经[%s]注册", identityInfo.Code, identityInfo.SecurityCode))
			return
		}
		// 没有被占用, 先创建空白用户
		user, e = domains.TxCreateSecurity(tx, platform)
		if e != nil {
			log.Info("创建用户失败")
			log.Info(e)
			return
		}
		// 绑定登录凭证
		identityInfo, e = domains.TxCreateIdentity(tx, user.Code, identity, identityType)
		if e != nil {
			log.Info(e)
			return
		}
		return
	})
	return
}
