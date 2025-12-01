package www

import (
	"time"

	"github.com/lishimeng/www/def"
	"github.com/lishimeng/www/dto"
)

var SuperIdentity dto.Identity

var SystemScope = def.App
var SystemMenuGroup def.MenuGroup
var AccessTokenTtl = time.Hour * 24 * 14
