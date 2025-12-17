package www

import (
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
)

var EnableAuthChk bool
var DisableAuthChkScope bool

// WithAuth token验证器,
// auth.JwtBasic header预处理
// auth.Forbidden401Handler 无权限时返回401, 返回格式按照参数 auth.ForbiddenOption
func WithAuth(handler func(ctx server.Context)) []server.Handler {
	var handlers []server.Handler
	if EnableAuthChk {
		var fbd []func(auth.ForbiddenOption) auth.ForbiddenOption
		fbd = append(fbd, auth.WithJsonResp())
		if !DisableAuthChkScope {
			fbd = append(fbd, auth.WithScope(SystemScope.Name))
		}
		handlers = append(handlers,
			auth.JwtBasic(),
			auth.Forbidden401Handler(fbd...))
	}
	handlers = append(handlers, handler)
	return handlers
}
