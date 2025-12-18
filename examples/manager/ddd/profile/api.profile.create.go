package profile

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
	"net/http"
)

func initializeApi(ctx server.Context) {
	var req dtoProfile
	var resp app.Response
	var err error

	userCode := ctx.C.Params().Get("code")
	if len(userCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "user code can not be null"
		ctx.Json(resp)
		return
	}

	err = ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	var profile usersTable.AuthUserProfile
	req.Transform(&profile)
	profile.Code = userCode
	_, err = app.GetOrm().Context.Insert(&profile)
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
