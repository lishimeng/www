package users

import (
	"net/http"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/x/container"
)

type RespMe struct {
	UserId   int      `json:"userId,omitempty"`
	Username string   `json:"username,omitempty"`
	Nickname string   `json:"nickname,omitempty"`
	Avatar   string   `json:"avatar,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Perms    []string `json:"perms,omitempty"`
}

func apiMeInfo(ctx server.Context) {
	//var err error
	var resp app.ResponseWrapper
	var data RespMe

	log.Info("get me info")

	data.UserId = 3
	data.Username = "admin"
	data.Nickname = "admin"
	data.Avatar = "https://foruda.gitee.com/images/1723603502796844527/03cdca2a_716974.gif"
	data.Roles = []string{"admin"}

	uid := auth.GetUid(ctx)
	if len(uid) > 0 {

		perms := handleMyPerms(uid)
		// 转换为列表
		for _, perm := range perms {
			data.Perms = append(data.Perms, perm.PermissionCode)
		}
	}

	if uid == www.SuperIdentity.SecurityCode {
		data.Perms = append(data.Perms,
			"sys:user:query",
			"sys:user:edit",
			"sys:menu:add",
			"sys:menu:edit",
			"sys:menu:delete",
		)
	}

	resp.Code = http.StatusOK
	resp.Data = data
	ctx.Json(resp)
}

func handleMyPerms(userCode string) (list []dto.AuthUserPerm) {
	var permManager www.PermManager
	_ = container.Get(&permManager)
	list = permManager.GetUserPerms(userCode)
	return
}
