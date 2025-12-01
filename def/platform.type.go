package def

type UserPlatform int
type Platform struct {
	Code UserPlatform
	Name string
}

const (
	AppUser     UserPlatform = 10
	SaasUser    UserPlatform = 20
	AdminUser   UserPlatform = 30
	ManagerUser UserPlatform = 40
)

var (
	App     = Platform{AppUser, "app"}
	Saas    = Platform{SaasUser, "saas"}
	Admin   = Platform{AdminUser, "admin"}
	Manager = Platform{ManagerUser, "manager"}
)
