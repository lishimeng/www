package roles

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
)

func apiDelRole(ctx server.Context) {
	var err error
	var resp app.ResponseWrapper
	roleCode := ctx.C.Params().GetStringDefault("code", "")
	if len(roleCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "role code can not be null"
		ctx.Json(resp)
		return
	}
	var m = getRoleOptManager()
	err = m.DelRole(roleCode)

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
