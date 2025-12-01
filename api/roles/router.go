package roles

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
	"github.com/lishimeng/x/container"
)

func getRoleManager() www.RoleManager {
	var m www.RoleManager
	_ = container.Get(&m)
	return m
}

func getRoleOptManager() www.RoleOperationManager {
	var m www.RoleOperationManager
	_ = container.Get(&m)
	return m
}

func Router(root server.Router) {

	r := root.Path("/role")     // 角色
	c := root.Path("/category") // 角色权限等级
	b := root.Path("/bind")     // 绑定用户, 角色:用户=1:n

	c.Get("/", apiRoleCategoryList) // role类型(级别)列表

	r.Get("/", www.WithAuth(apiRoleList)...) // role列表
	r.Get("/{roleCode}", apiGetRole)         // 读取role详情
	r.Post("/", www.WithAuth(apiAddRole)...) // 添加role
	r.Patch("/", www.WithAuth(apiUpdateRole)...)
	r.Delete("/{code}", www.WithAuth(apiDelRole)...) // 删除role

	// -----------绑定关系
	b.Post("/", www.WithAuth(apiUserBindRole)...)                  // 用户绑定角色(添加user与role的关系)
	b.Delete("/", www.WithAuth(apiUserUnbindRole)...)              // 用户解绑角色(删除user与role的关系)
	b.Get("/{roleCode}", www.WithAuth(apiUserBindUserRoleList)...) // 角色下绑定的用户列表
}
