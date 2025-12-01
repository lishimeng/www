package dto

type RoleCategory struct { // role category只管理级别, 而且不做数据结构上的管控, 由各个系统的业务逻辑自行管理
	Id   int    `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type Role struct {
	Id               int    `json:"id,omitempty"`
	Code             string `json:"code,omitempty"`
	Name             string `json:"name,omitempty"`
	RoleCategoryCode string `json:"roleCategoryCode,omitempty"`
	RoleValue        string `json:"roleValue,omitempty"`
	ProjectCode      string `json:"projectCode,omitempty"`
	IsGlobal         bool   `json:"isGlobal,omitempty"`
}

type UserRole struct {
	Id               int    `json:"id,omitempty"`
	UserCode         string `json:"userCode,omitempty"`
	RoleCategoryCode string `json:"roleCategoryCode,omitempty"`
	RoleCode         string `json:"roleCode,omitempty"`
	RoleValue        string `json:"roleValue,omitempty"`
	ProjectCode      string `json:"projectCode,omitempty"`
	IsGlobal         bool   `json:"isGlobal,omitempty"`
}
