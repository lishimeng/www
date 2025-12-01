package roles

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
)

func apiRoleList(ctx server.Context) {
	var resp app.ResponseWrapper
	var m = getRoleOptManager()

	list, err := m.All()
	if err != nil {
		log.Info(err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	resp.Code = http.StatusOK
	resp.Data = list // 全量数据,不需要分页
	ctx.Json(resp)
}
