package users

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/x/container"
)

func apiSaasUserList(ctx server.Context) {

	var err error
	var resp app.ResponseWrapper
	pageNum, pageSize := def.GetParamPager(ctx)

	filterEnterprise := ctx.C.URLParamDefault("enterprise", "")

	var saasOptMgr www.SaasOperationManager
	_ = container.Get(&saasOptMgr)
	pager := app.Pager[dto.SaasUser]{}
	pager.PageSize = pageSize
	pager.PageNum = pageNum
	err = saasOptMgr.SaasList(filterEnterprise, &pager)
	if err != nil {
		resp.Code = 40010
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	resp.Code = http.StatusOK
	resp.Data = pager
	resp.Message = "success"
	ctx.Json(resp)
}
