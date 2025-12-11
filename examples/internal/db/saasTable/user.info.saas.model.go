package saasTable

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
)

type SaasUserInfo struct {
	// 企业用户
	app.Pk
	Code           string `orm:"column(code);unique"`
	Name           string `orm:"column(name);null"`
	EnterpriseCode string `orm:"column(enterprise_code);null"`
	SecurityCode   string `orm:"column(security_code)"` // user_code+security_code唯一
	app.TableChangeInfo
}

func (s *SaasUserInfo) Transform(dst *dto.SaasUser) {
	dst.Id = s.Id
	dst.EnterpriseCode = s.EnterpriseCode
	dst.SecurityCode = s.SecurityCode
	dst.Name = s.Name
	dst.Code = s.Code
	dst.Status = def.SaasUserStatus(s.Status)
}

func (s *SaasUserInfo) Convert(src dto.SaasUser) {
	s.EnterpriseCode = src.EnterpriseCode
	s.SecurityCode = src.SecurityCode
	s.Name = src.Name
	s.Code = src.Code
}
