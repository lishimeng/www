package dto

// Permission 权限: 通过 Category 按照 Platform 分级
// 全局配置，不关联特定的 Saas project 或 organization
type Permission struct {
	Id          int    `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
}

type AuthUserPerm struct {
	Id             int    `json:"id,omitempty"`
	RoleCode       string `json:"roleCode,omitempty"`
	SaasUserCode   string `json:"saasUserCode,omitempty"`
	PermissionCode string `json:"permissionCode,omitempty"`
	IsGlobal       int    `json:"isGlobal,omitempty"`
}

type AuthRolePerm struct {
	Id             int    `json:"id,omitempty"`
	RoleCode       string `json:"roleCode,omitempty"`
	PermissionCode string `json:"permissionCode,omitempty"`
}
