package auth

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"net/http"
)

func logoutApi(ctx server.Context) {
	var resp app.Response
	// todo: 禁用 refresh code
	resp.Code = http.StatusOK
	resp.Message = "success"
	ctx.Json(resp)
}
