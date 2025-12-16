package roles

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
)

type FormUserRole struct {
	SaasUserCode string `json:"saasUserCode,omitempty"` // Saas user 的唯一码
	RoleCode     string `json:"roleCode,omitempty"`
	ProjectCode  string `json:"projectCode,omitempty"`
}

func apiUserBindRole(ctx server.Context) {
	var req FormUserRole
	var err error
	var resp app.ResponseWrapper

	err = ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	if len(req.RoleCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "role code can not be null"
		ctx.Json(resp)
		return
	}
	if len(req.SaasUserCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "user code can not be null"
		ctx.Json(resp)
		return
	}
	var m = getRoleOptManager()
	_, err = m.BindUser(req.SaasUserCode, req.RoleCode)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "success"
	ctx.Json(resp)
}

func apiUserUnbindRole(ctx server.Context) {
	var req FormUserRole
	var err error
	var resp app.ResponseWrapper
	err = ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	if len(req.SaasUserCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "user code can not be null"
		ctx.Json(resp)
	}
	if len(req.RoleCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "role code can not be null"
		ctx.Json(resp)
		return
	}
	var m = getRoleOptManager()
	err = m.UnbindUser(req.SaasUserCode, req.RoleCode)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "success"
	ctx.Json(resp)
}

func apiUserBindUserRoleList(ctx server.Context) {
	var resp app.ResponseWrapper
	roleCode := ctx.C.Params().GetStringDefault("roleCode", "")
	if len(roleCode) == 0 {
		resp.Code = http.StatusOK
		resp.Data = []string{}
		ctx.Json(resp)
		return
	}
	var m = getRoleOptManager()
	urs, err := m.RoleUsers(roleCode)
	if err != nil {
		log.Info(err)
		resp.Code = http.StatusOK
		resp.Data = []string{}
		ctx.Json(resp)
		return
	}
	resp.Code = http.StatusOK
	resp.Data = urs
	ctx.Json(resp)
}
