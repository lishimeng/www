package permissions

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/www"
)

func Router(root server.Router) {
	// 权限管理
	p := root.Path("/permission")       // 权限
	rp := root.Path("/role-permission") // 角色权限绑定

	p.Get("/", www.WithAuth(apiPermissionList)...)                // 权限列表
	p.Get("/{permissionCode}", www.WithAuth(apiGetPermission)...) // 获取单个权限
	p.Post("/", www.WithAuth(apiAddPermission)...)                // 添加权限
	p.Patch("/", www.WithAuth(apiUpdatePermission)...)            // 更新权限
	p.Delete("/{code}", www.WithAuth(apiDelPermission)...)        // 删除权限
	rp.Get("/{roleCode}", www.WithAuth(apiRolePermissions)...)    // 获取角色的权限列表
	rp.Post("/", www.WithAuth(apiRoleBindPermission)...)          // 绑定角色权限
	rp.Delete("/", www.WithAuth(apiRoleUnbindPermission)...)      // 解绑角色权限
}
