package domain

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/examples/internal/db/saasTable"
)

func (h handler) CreateSaasUser(userCode string, orgCode string) (u dto.SaasUser, err error) {
	var model saasTable.SaasUserInfo
	model.Code = userCode
	model.EnterpriseCode = orgCode
	// todo: 查询用户信息
	_, err = app.GetOrm().Context.Insert(&model)
	return
}

func (h handler) SaasUserPage(orgCode string, pager *app.Pager[dto.SaasUser]) (err error) {
	var simplePager app.SimplePager[saasTable.SaasUserInfo, dto.SaasUser]
	simplePager.Pager = *pager
	simplePager.Transform = func(src saasTable.SaasUserInfo, dst *dto.SaasUser) {
		src.Transform(dst)
	}
	simplePager.QueryBuilder = func(tx persistence.TxContext) any {
		q := tx.Context.QueryTable(new(saasTable.SaasUserInfo))
		if len(orgCode) > 0 {
			q = q.Filter("EnterpriseCode", orgCode)
		}
		return q
	}
	err = app.QueryPage(&simplePager)
	*pager = simplePager.Pager
	return err
}
