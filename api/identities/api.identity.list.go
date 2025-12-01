package identities

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/x/container"
)

func apiIdentityList(ctx server.Context) {
	var resp app.ResponseWrapper
	var err error

	var identityManager www.IdentityManager
	_ = container.Get(&identityManager)

	pageNum, pageSize := def.GetParamPager(ctx)
	pager := app.Pager[dto.Identity]{}
	pager.PageSize = pageSize
	pager.PageNum = pageNum
	err = identityManager.Page(&pager)

	if err != nil {
		log.Info("获取identity列表失败: %v", err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Data = pager
	ctx.Json(resp)
}
