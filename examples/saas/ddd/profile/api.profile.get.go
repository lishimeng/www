package profile

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
	"net/http"
)

// 可自行配置的个人信息
type respProfile struct {
	UserId     int    `json:"userId,omitempty"`
	Username   string `json:"username,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	CreateTime string `json:"createTime,omitempty"` // 精确到日期
	// todo: identities...
}

func getProfileApi(ctx server.Context) {
	var resp app.ResponseWrapper
	var err error
	var data respProfile

	uid := auth.GetUid(ctx)
	if len(uid) == 0 {
		resp.Code = http.StatusBadRequest
		resp.Message = "user code can not be null"
		ctx.Json(resp)
		return
	}

	var profile usersTable.AuthUserProfile
	profile.Code = uid
	err = app.GetOrm().Context.Read(&profile, "Code")
	if err != nil {
		resp.Code = http.StatusNotFound
		resp.Message = "not found"
		ctx.Json(resp)
		return
	}

	data.UserId = profile.Id
	data.Username = profile.Name
	data.Nickname = profile.Name
	data.Avatar = profile.Avatar

	resp.Code = http.StatusOK
	resp.Data = data
	ctx.Json(resp)
}
