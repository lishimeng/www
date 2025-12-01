package dto

type SaasUser struct {
	Id             int    `json:"id,omitempty"`
	Code           string `json:"code,omitempty"`
	Name           string `json:"name,omitempty"`
	EnterpriseCode string `json:"enterpriseCode,omitempty"`
	SecurityCode   string `json:"securityCode,omitempty"`
}
