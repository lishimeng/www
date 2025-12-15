package viewToken

type DtoViewAccessToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"` // Bearer
	ExpiresIn    int64  `json:"expiresIn"`
}
