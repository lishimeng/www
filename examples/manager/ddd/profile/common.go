package profile

import "github.com/lishimeng/www/examples/internal/db/auth/usersTable"

type dtoProfile struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func getDefaultAvatar() string {
	return "https://foruda.gitee.com/images/1723603502796844527/03cdca2a_716974.gif"
}

func (p *dtoProfile) Transform(dst *usersTable.AuthUserProfile) {
	dst.Name = p.Name
	dst.Avatar = p.Avatar
}

func (p *dtoProfile) Convert(src usersTable.AuthUserProfile) {
	p.Name = src.Name
	p.Avatar = src.Avatar
	if len(p.Avatar) == 0 {
		p.Avatar = getDefaultAvatar()
	}
}
