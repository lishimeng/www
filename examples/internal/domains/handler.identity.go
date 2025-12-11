package domains

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
	"github.com/lishimeng/x/util"
	"regexp"
	"strings"
)

func TxGetIdentity(tx persistence.TxContext, identityCode string) (one dto.Identity, err error) {
	var model usersTable.AuthIdentity
	model.Code = identityCode
	if err = tx.Context.Read(&model, "Code"); err != nil {
		return
	}
	model.Transform(&one)
	return
}

func TxGetSecurity(tx persistence.TxContext, securityUserCode string) (model usersTable.AuthSecurityInfo, err error) {
	model.Code = securityUserCode
	if err = tx.Context.Read(&model, "Code"); err != nil {
		return
	}
	return
}

func TxCreateSecurity(tx persistence.TxContext, platform def.UserPlatform) (user dto.SecurityInfo, err error) {
	var model usersTable.AuthSecurityInfo
	model.Platform = platform
	model.Code = util.UUIDString()
	if _, err = tx.Context.Insert(&model); err != nil {
		return
	}
	model.Transform(&user)
	return
}

func TxCreateIdentity(tx persistence.TxContext, securityUserCode, identityCode string, identityType def.IdentityType) (one dto.Identity, err error) {
	var model usersTable.AuthIdentity
	model.SecurityCode = securityUserCode
	model.Code = identityCode
	model.IdentityType = identityType
	if _, err = tx.Context.Insert(&model); err != nil {
		return
	}
	model.Transform(&one)
	return
}

func TxUpdateSecurity(tx persistence.TxContext, model usersTable.AuthSecurityInfo, fields ...string) (err error) {
	_, err = tx.Context.Update(&model, fields...)
	return err
}

func (h handler) TxGetIdentity(tx persistence.TxContext, identityCode string) (one dto.Identity, err error) {
	return TxGetIdentity(tx, identityCode)
}

func (h handler) TxGetSecurity(tx persistence.TxContext, securityUserCode string) (one dto.SecurityInfo, err error) {
	var model usersTable.AuthSecurityInfo
	model, err = TxGetSecurity(tx, securityUserCode)
	if err != nil {
		return
	}
	model.Transform(&one)
	return
}

func (h handler) TxCreateIdentity(tx persistence.TxContext, securityUserCode, identityCode string, identityType def.IdentityType) (one dto.Identity, err error) {
	return TxCreateIdentity(tx, securityUserCode, identityCode, identityType)
}

func (h handler) TxDelIdentity(tx persistence.TxContext, code string) (err error) {
	_, err = tx.Context.QueryTable(new(usersTable.AuthIdentity)).
		Filter("Code", code).
		Delete()
	return
}

func (h handler) TxGetIdentities(tx persistence.TxContext, securityUserFilter string, limits ...int) (data []dto.Identity, err error) {
	var models []usersTable.AuthIdentity
	q := tx.Context.QueryTable(new(usersTable.AuthIdentity)).
		Filter("SecurityCode", securityUserFilter)
	if len(limits) > 0 {
		q = q.Limit(limits[0])
	}
	_, err = q.All(&models)
	if err != nil {
		return
	}
	for _, model := range models {
		var one dto.Identity
		model.Transform(&one)
		data = append(data, one)
	}
	return
}

func (h handler) GetIdentityPage(pager *app.Pager[dto.Identity]) (err error) {
	var simplePager app.SimplePager[usersTable.AuthIdentity, dto.Identity]
	simplePager.Pager = *pager
	simplePager.Transform = func(src usersTable.AuthIdentity, dst *dto.Identity) {
		src.Transform(dst)
	}
	simplePager.QueryBuilder = func(tx persistence.TxContext) any {
		q := tx.Context.QueryTable(new(usersTable.AuthIdentity))
		return q
	}
	err = app.QueryPage(&simplePager)
	*pager = simplePager.Pager
	return err
}

func (h handler) GetSecurityInfo(code string) (one dto.SecurityInfo, err error) {
	var model usersTable.AuthSecurityInfo
	model.Code = code
	if err = app.GetOrm().Context.Read(&model, "Code"); err != nil {
		return
	}
	model.Transform(&one)
	return
}

func (h handler) DetectIdentityType(identityCode string) def.IdentityType {
	// 判断是否为邮箱（包含 @ 符号）
	if strings.Contains(identityCode, "@") {
		return def.LoginEmail
	}

	// 判断是否为手机号（11位数字，以1开头）
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if phoneRegex.MatchString(identityCode) {
		return def.LoginPhone
	}

	// 判断是否为微信相关（包含特定前缀或格式）
	if strings.HasPrefix(identityCode, "wx_") || strings.HasPrefix(identityCode, "unionid_") {
		return def.LoginWechatUnionId
	}
	if strings.HasPrefix(identityCode, "openid_") {
		return def.LoginWechatOpenId
	}

	// 默认使用登录码
	return def.LoginCode
}
