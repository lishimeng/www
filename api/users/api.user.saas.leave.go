package users

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/x/container"
	"net/http"
)

type ReqSaasLeave struct {
	UserCode string `json:"userCode"`
	OrgCode  string `json:"orgCode"`
}

func apiSaasLeave(ctx server.Context) {

	var err error
	var req ReqSaasJoin
	var resp app.Response

	err = ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	if len(req.UserCode) == 0 || len(req.OrgCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "userCode or orgCode is empty"
		ctx.Json(resp)
		return
	}

	var saasOptMgr www.SaasOperationManager
	_ = container.Get(&saasOptMgr)
	err = saasOptMgr.LeaveOrganization(req.OrgCode, req.UserCode)

	if err != nil {
		log.Info("退出saas组织失败")
		log.Info(err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp.Code = http.StatusOK
	ctx.Json(resp)
}
