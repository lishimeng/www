package domains

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/examples/internal/db/saasTable"
)

func (h handler) TxCreateSaasOrg(tx persistence.TxContext, orgCode string, name string, admin string) (org dto.SaasOrganization, err error) {
	var id int64
	var model saasTable.SaasOrganization
	model.Name = name
	model.Code = orgCode
	model.AdminUser = admin
	id, err = tx.Context.Insert(&model)
	model.Id = int(id)
	model.Transform(&org)
	return
}

func (h handler) TxAddOrgUser(tx persistence.TxContext, dto *dto.SaasUser) (err error) {
	var model saasTable.SaasUserInfo
	model.Code = dto.Code
	err = tx.Context.Read(&model, "Code")
	if err == nil {
		model.Convert(*dto)
		model.Status = int(def.SaasUserActive)
		_, err = tx.Context.Update(&model)
		return
	}
	model.Convert(*dto)
	model.Status = int(def.SaasUserActive)
	_, err = tx.Context.Insert(&model)
	return
}

func (h handler) TxGetOrg(tx persistence.TxContext, orgCode string) (org dto.SaasOrganization, err error) {
	var model saasTable.SaasOrganization
	model.Code = orgCode
	err = tx.Context.Read(&model, "Code")
	if err != nil {
		return
	}
	model.Transform(&org)
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
func (h handler) TxDelOrgUser(tx persistence.TxContext, code string) (err error) {
	_, err = tx.Context.QueryTable(new(saasTable.SaasUserInfo)).
		Filter("Code", code).
		Update(map[string]any{"Status": def.SaasUserDisabled})
	return
}

func (h handler) TxGetSaasUser(tx persistence.TxContext, code string) (u dto.SaasUser, err error) {
	var model saasTable.SaasUserInfo
	err = tx.Context.QueryTable(new(saasTable.SaasUserInfo)).
		Filter("Code", code).
		Filter("Status", def.SaasUserActive).
		One(&model)
	if err != nil {
		return
	}
	model.Transform(&u)
	return
}

func (h handler) GetAllOrgs() (list []dto.SaasOrganization, err error) {
	var model []saasTable.SaasOrganization
	_, err = app.GetOrm().Context.QueryTable(new(saasTable.SaasOrganization)).
		All(&model)
	if err != nil {
		return
	}
	for _, v := range model {
		var u dto.SaasOrganization
		v.Transform(&u)
		list = append(list, u)
	}
	return
}

func (h handler) GetSaasOrg(code string) (one dto.SaasOrganization, err error) {
	var model saasTable.SaasOrganization
	model.Code = code
	err = app.GetOrm().Context.Read(&model, "Code")
	if err != nil {
		return
	}
	model.Transform(&one)
	return
}

func (h handler) GetUserSaasOrgs(userCode string) (list []dto.SaasUser, err error) {
	var model []saasTable.SaasUserInfo
	_, err = app.GetOrm().Context.QueryTable(new(saasTable.SaasUserInfo)).
		Filter("SecurityCode", userCode).
		All(&model)
	if err != nil {
		return
	}
	for _, v := range model {
		var u dto.SaasUser
		v.Transform(&u)
		list = append(list, u)
	}
	return
}
