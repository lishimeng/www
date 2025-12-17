package security

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
	"net/http"
)

type reqResetPassword struct {
	SecurityCode string `json:"securityCode"`
	NewPassword  string `json:"newPassword"`
}

// 特权api，绕开旧密码校验
func resetPasswordApi(ctx server.Context) {
	var req reqResetPassword
	var err error
	var resp app.Response

	err = ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	if len(req.SecurityCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "identity can not be null"
		ctx.Json(resp)
		return
	}

	if len(req.NewPassword) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "password can not be null"
		ctx.Json(resp)
		return
	}

	// 设置密码
	err = UpdatePsw(req.SecurityCode, req.NewPassword, func(_ usersTable.AuthSecurityInfo) bool {
		return true
	})
	if err != nil {
		log.Info("设置密码失败: %v", err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp.Code = http.StatusOK
	ctx.Json(resp)
}
