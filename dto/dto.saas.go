package dto

import "github.com/lishimeng/www/def"

type SaasUser struct {
	Id             int                `json:"id,omitempty"`
	Code           string             `json:"code,omitempty"`
	Name           string             `json:"name,omitempty"`
	EnterpriseCode string             `json:"enterpriseCode,omitempty"`
	SecurityCode   string             `json:"securityCode,omitempty"`
	Status         def.SaasUserStatus `json:"status,omitempty"` // 是否已禁用
}

type SaasOrganization struct {
	Id    int    `json:"id,omitempty"`
	Code  string `json:"code,omitempty"`
	Name  string `json:"name,omitempty"`
	Admin string `json:"admin,omitempty"`
}
