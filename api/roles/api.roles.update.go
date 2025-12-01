package roles

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www/dto"
)

func apiUpdateRole(ctx server.Context) {
	var err error
	var roleForm dto.Role
	var resp app.ResponseWrapper

	err = ctx.C.ReadJSON(&roleForm)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	var m = getRoleOptManager()
	err = m.EditRole(&roleForm)

	if err != nil {
		log.Info("更新角色失败")
		log.Info(err)
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	resp.Code = http.StatusOK
	resp.Data = roleForm
	ctx.Json(resp)
}
