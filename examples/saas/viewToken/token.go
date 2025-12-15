package viewToken

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"time"

	"gitee.com/xdho/kcbz/internal/ddd/infrastructures"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/x/container"
)

const (
	ttl = time.Hour * 6
)

func GenRefreshToken(accessToken string) (token string) {
	tmp := fmt.Sprintf("%s%d%s", accessToken, time.Now().Unix(), infrastructures.GenCode())
	bs := sha512.New().Sum([]byte(tmp))
	token = base64.URLEncoding.EncodeToString(bs[:])
	return
}

func GenToken(uCode, loginType, scope string, ttl time.Duration) (tokenVal []byte, err error) {
	var tokenPayload token.JwtPayload
	tokenPayload.Uid = uCode
	tokenPayload.Client = loginType
	tokenPayload.Scope = scope
	tokenVal, err = _genWithTTL(tokenPayload, ttl)
	return
}

// GenOpenToken SaaS服务，供第三方使用
func GenOpenToken(userCode string, appID string, org string, scope string) (tokenVal []byte, err error) {
	var tokenPayload token.JwtPayload
	tokenPayload.Uid = userCode
	tokenPayload.Client = appID
	tokenPayload.Org = org
	tokenPayload.Scope = scope
	// 所有access token使用统一ttl，无需传入参数
	tokenVal, err = _genWithTTL(tokenPayload, ttl)
	return
}

func _genToken(payload token.JwtPayload) (content []byte, err error) {
	var provider token.JwtProvider
	err = container.Get(&provider)
	if err != nil {
		return
	}
	log.Info("tokenPayload: %s", payload)
	content, err = provider.Gen(payload)
	return
}

func _genWithTTL(payload token.JwtPayload, ttl time.Duration) (content []byte, err error) {
	var provider token.JwtProvider
	err = container.Get(&provider)
	if err != nil {
		return
	}
	log.Info("tokenPayload: %s", payload)
	content, err = provider.GenWithTTL(payload, ttl)
	return
}
