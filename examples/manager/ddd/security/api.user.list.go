package security

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
	"net/http"
)

func getListApi(ctx server.Context) {
	var resp app.ResponseWrapper
	var err error

	pageNum, pageSize := def.GetParamPager(ctx)
	pager := app.Pager[dto.SecurityInfo]{}
	pager.PageSize = pageSize
	pager.PageNum = pageNum

	var simplePager app.SimplePager[usersTable.AuthSecurityInfo, dto.SecurityInfo]
	simplePager.Pager = pager
	simplePager.Transform = func(src usersTable.AuthSecurityInfo, dst *dto.SecurityInfo) {
		src.Transform(dst)
	}
	simplePager.QueryBuilder = func(tx persistence.TxContext) any {
		q := tx.Context.QueryTable(new(usersTable.AuthSecurityInfo))
		return q
	}
	simplePager.OrderByExp = append(simplePager.OrderByExp, "-CreateTime")
	err = app.QueryPage(&simplePager)
	pager = simplePager.Pager

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
