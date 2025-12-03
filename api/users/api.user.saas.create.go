package users

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/x/container"
)

type ReqSaasCreate struct {
	UserCode string `json:"userCode"`
	OrgCode  string `json:"orgCode"`
}

// Notice: 为一个securityUser添加一个saas用户(企业账号认证, 共享一套登录凭证), 而不是从头创建一个新账号
// 认证成企业，并成为企业的超级管理员
func apiCreateSaasUser(ctx server.Context) {

	var err error
	var req ReqSaasCreate
	var resp app.ResponseWrapper

	err = ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	if len(req.UserCode) == 0 || len(req.UserCode) == 0 || len(req.OrgCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "phone or email or password is empty"
		ctx.Json(resp)
		return
	}
	var saasOptMgr www.SaasOperationManager
	_ = container.Get(&saasOptMgr)
	u, err := saasOptMgr.CreateSaas(req.UserCode, req.OrgCode)

	if err != nil {
		log.Info("创建saas用户失败")
		log.Info(err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Data = u
	ctx.Json(resp)
}
