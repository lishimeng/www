package profile

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
	"github.com/lishimeng/www/examples/internal/db/saasTable"
	"net/http"
)

func updateApi(ctx server.Context) {
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
	profile.Code = userCode
	err = app.GetOrm().Context.Read(&profile, "Code")
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	cols := make([]string, 0)
	if len(req.Name) > 0 && req.Name != profile.Name {
		profile.Name = req.Name
		cols = append(cols, "Name")

		var saasUser saasTable.SaasUserInfo
		saasUser.Code = userCode
		err = app.GetOrm().Context.Read(&saasUser, "Code")
		if err == nil {
			saasUser.Name = req.Name
			_, err = app.GetOrm().Context.Update(&saasUser, "Name")
		}
	}
	if len(req.Avatar) > 0 {
		profile.Avatar = req.Avatar
		cols = append(cols, "Avatar")
	}
	_, err = app.GetOrm().Context.Update(&profile, cols...)
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
