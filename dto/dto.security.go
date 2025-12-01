package dto

type SecurityInfo struct {
	Id             int    `json:"id"`
	Code           string `json:"code"`           // 用户编号
	Platform       int    `json:"platform"`       // 用户平台 10-AppUser, 20-SaasUser, 30-AdminUser
	SecondFaEnable int    `json:"secondFaEnable"` // 是否启用2FA
	SecondFaCode   string `json:"secondFaCode"`   // 2FA设备编号
	Status         int    `json:"status"`
	CreateTime     string `json:"createTime"`
	UpdateTime     string `json:"updateTime"`
}
