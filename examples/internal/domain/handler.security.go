package domain

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/www/dto"
	"github.com/lishimeng/www/examples/internal/db/auth/usersTable"
)

func (h handler) GetSecurityInfo(code string) (one dto.SecurityInfo, err error) {
	var model usersTable.AuthSecurityInfo
	model.Code = code
	if err = app.GetOrm().Context.Read(&model, "Code"); err != nil {
		return
	}
	model.Transform(&one)
	return
}
