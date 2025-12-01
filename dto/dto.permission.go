package dto

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
	UserCode       string `json:"userCode,omitempty"`
	PermissionCode string `json:"permissionCode,omitempty"`
	IsGlobal       int    `json:"isGlobal,omitempty"`
}

type AuthRolePerm struct {
	Id             int    `json:"id,omitempty"`
	RoleCode       string `json:"roleCode,omitempty"`
	PermissionCode string `json:"permissionCode,omitempty"`
}
