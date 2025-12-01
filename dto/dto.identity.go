package dto

type Identity struct {
	Id           int    `json:"id"`
	Code         string `json:"code"`         // 登录标识
	IdentityType int    `json:"identityType"` // 类型
	SecurityCode string `json:"securityCode"` // 用户码
	Status       int    `json:"status"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}
