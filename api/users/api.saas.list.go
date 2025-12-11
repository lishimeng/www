package users

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"net/http"
)

func apiSaasList(ctx server.Context) {
	var resp app.ResponseWrapper
	var m = getSaasOptManager()

	list, err := m.AllOrgs()
	if err != nil {
		log.Info(err)
		resp.Code = http.StatusNotFound
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	resp.Code = http.StatusOK
	resp.Data = list // 全量数据,不需要分页
	ctx.Json(resp)
}

func apiSassOne(ctx server.Context) {
	var resp app.ResponseWrapper
	var m = getSaasManager()

	code := ctx.C.Params().Get("orgCode")
	list, err := m.SaasOrg(code)
	if err != nil {
		log.Info(err)
		resp.Code = http.StatusNotFound
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Data = list
	ctx.Json(resp)
}
