package security

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/x/container"
	"net/http"
)

func getIdentityApi(ctx server.Context) {
	var resp app.ResponseWrapper

	userCode := ctx.C.Params().Get("code")

	var identityManager www.IdentityManager
	_ = container.Get(&identityManager)

	identities, err := identityManager.UserIdentities(userCode)

	if err != nil {
		log.Info("获取identity列表失败: %v", err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Data = identities
	ctx.Json(resp)
}
