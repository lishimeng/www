package profile

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
	"net/http"
)

func getOneApi(ctx server.Context) {
	var dto dtoProfile
	var resp app.ResponseWrapper
	var err error

	userCode := ctx.C.Params().Get("code")
	if len(userCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "user code can not be null"
		ctx.Json(resp)
		return
	}

	var profile usersTable.AuthUserProfile
	profile.Code = userCode
	err = app.GetOrm().Context.Read(&profile, "Code")
	if err != nil {
		resp.Code = http.StatusNotFound
		resp.Message = "not found"
		ctx.Json(resp)
		return
	}

	dto.Convert(profile)
	resp.Code = http.StatusOK
	resp.Message = "success"
	resp.Data = dto
	ctx.Json(resp)
}
