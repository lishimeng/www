package www

import (
	"errors"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
)

type DbIdentityManager struct {
	domain DomainHandler
}

func NewDbIdentityManager(handler DomainHandler) IdentityManager {
	var h IdentityManager
	h = &DbIdentityManager{
		domain: handler,
	}
	return h
}

func (im *DbIdentityManager) UserIdentities(userCode string) (list []dto.Identity, err error) {
	err = app.GetOrm().Transaction(func(tx persistence.TxContext) (e error) {
		list, e = im.domain.TxGetIdentities(tx, userCode, 100)
		return
	})
	return
}

func (im *DbIdentityManager) Page(pager *app.Pager[dto.Identity]) (err error) {
	err = im.domain.GetIdentityPage(pager)
	return
}

func (im *DbIdentityManager) Bind(identityCode, securityUserCode string, identityType def.IdentityType) (err error) {
	err = app.Transaction(func(tx persistence.TxContext) (e error) {
		// 检查 auth_identity 是否在系统中已经存在, 如果存在, 不能绑定
		_, e = im.domain.TxGetIdentity(tx, identityCode)
		if e == nil {
			log.Info("登录凭证已存在: %s", identityCode)
			e = errors.New("登录凭证已存在: " + identityCode)
			return
		}

		// 检查 security_user 是否在系统中存在, 如果不存在, 不能绑定
		_, e = im.domain.TxGetSecurity(tx, securityUserCode)
		if e != nil {
			log.Info("用户不存在: %s", securityUserCode)
			e = errors.New("用户不存在: " + securityUserCode)
			return
		}

		// 执行创建 auth_identity
		_, e = im.domain.TxCreateIdentity(tx, securityUserCode, identityCode, identityType)
		if e != nil {
			log.Info("创建登录凭证失败: %s", identityCode)
			return
		}

		log.Info("绑定登录凭证成功: [%s] -> [%s]", identityCode, securityUserCode)
		return nil
	})
	return
}
func (im *DbIdentityManager) Unbind(identityCode string, securityUserCode string) (err error) {
	err = app.Transaction(func(tx persistence.TxContext) (e error) {
		// 检查 identity 存在
		identity, e := im.domain.TxGetIdentity(tx, identityCode)
		if e != nil {
			log.Info("登录凭证不存在: %s", identityCode)
			e = errors.New("登录凭证不存在: " + identityCode)
			return
		}

		// 检查 identity 绑定的 security_user 是不是参数中的 securityUserCode, 如果不匹配,不能删除
		if identity.SecurityCode != securityUserCode {
			log.Info("登录凭证不属于该用户: [%s] -> [%s] (期望: %s)", identityCode, identity.SecurityCode, securityUserCode)
			e = errors.New("登录凭证不属于该用户")
			return
		}

		// 检查当前 security_user 绑定的 identity 数量, 如果仅有一个, 不能删除
		identities, e := im.domain.TxGetIdentities(tx, securityUserCode, 5)
		if e != nil {
			log.Info("查询用户登录凭证列表失败: %s", securityUserCode)
			return
		}

		if len(identities) <= 1 {
			log.Info("用户仅有一个登录凭证, 不能删除: %s", securityUserCode)
			e = errors.New("用户仅有一个登录凭证, 不能删除")
			return
		}

		// 执行删除
		e = im.domain.TxDelIdentity(tx, identityCode)
		if e != nil {
			log.Info("删除登录凭证失败: %s", identityCode)
			return
		}

		log.Info("删除登录凭证成功: [%s] <-x-> [%s]", identityCode, securityUserCode)
		return nil
	})
	return
}

func (im *DbIdentityManager) DetectIdentityType(identityCode string) def.IdentityType {
	return im.domain.DetectIdentityType(identityCode)
}
