package profile

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
	"github.com/lishimeng/x/container"
	"net/http"
)

type RespMe struct {
	UserId   int      `json:"userId,omitempty"`
	Username string   `json:"username,omitempty"`
	Nickname string   `json:"nickname,omitempty"`
	Avatar   string   `json:"avatar,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Perms    []string `json:"perms,omitempty"`
}

func getMeApi(ctx server.Context) {
	var resp app.ResponseWrapper
	var err error
	var data RespMe

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

	perms := handleMyPerms(uid)
	// 转换为列表
	for _, perm := range perms {
		data.Perms = append(data.Perms, perm.PermissionCode)
	}

	roles, _ := handleMyRoles(uid)
	for _, role := range roles {
		data.Roles = append(data.Roles, role.RoleCode)
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

func handleMyRoles(userCode string) (list []dto.UserRole, err error) {
	var roleManager www.RoleManager
	_ = container.Get(&roleManager)
	list, err = roleManager.UserRoles(userCode)
	return
}
