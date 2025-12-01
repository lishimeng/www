package users

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/x/container"
)

type UserDetailResponse struct {
	SecurityInfo *dto.SecurityInfo `json:"securityInfo,omitempty"`
	Identities   []dto.Identity    `json:"identities"`
}

func apiUserDetail(ctx server.Context) {
	var resp app.ResponseWrapper
	securityCode := ctx.C.Params().GetStringDefault("securityCode", "")

	if len(securityCode) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "security code can not be null"
		ctx.Json(resp)
		return
	}

	var detail UserDetailResponse

	aom := getAuthOptManager()
	securityInfo, err := aom.SecurityUser(securityCode)
	if err != nil {
		log.Info("获取用户详情失败: %v", err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}

	var identityMgr www.IdentityManager
	_ = container.Get(&identityMgr)
	identities, _ := identityMgr.UserIdentities(securityInfo.Code)
	detail.Identities = identities

	resp.Code = http.StatusOK
	resp.Data = detail
	ctx.Json(resp)
}
