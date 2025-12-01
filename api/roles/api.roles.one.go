package roles

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
)

func apiGetRole(ctx server.Context) {
	var err error
	var resp app.ResponseWrapper

	roleCode := ctx.C.Params().GetStringDefault("roleCode", "")
	if len(roleCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "role code is required"
		ctx.Json(resp)
		return
	}
	var m = getRoleOptManager()
	dto, err := m.One(roleCode)

	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	resp.Code = http.StatusOK
	resp.Data = dto
	ctx.Json(resp)
}
