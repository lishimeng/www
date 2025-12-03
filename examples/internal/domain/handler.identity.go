package domain

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
)

func (h handler) TxGetIdentity(tx persistence.TxContext, identityCode string) (one dto.Identity, err error) {
	var model usersTable.AuthIdentity
	model.Code = identityCode
	if err = tx.Context.Read(&model, "Code"); err != nil {
		return
	}
	model.Transform(&one)
	return
}

func (h handler) TxGetSecurity(tx persistence.TxContext, securityUserCode string) (one dto.SecurityInfo, err error) {
	var model usersTable.AuthSecurityInfo
	model.Code = securityUserCode
	if err = tx.Context.Read(&model, "Code"); err != nil {
		return
	}
	model.Transform(&one)
	return
}

func (h handler) TxCreateIdentity(tx persistence.TxContext, securityUserCode, identityCode string, identityType def.IdentityType) (one dto.Identity, err error) {
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
