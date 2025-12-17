package auth

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www/examples/saas/storage"
	"github.com/lishimeng/www/examples/saas/viewToken"
	"net/http"
	"time"
)

const captchaPrefix = "captcha:"

type respCaptcha struct {
	CaptchaKey    string `json:"captchaKey"`
	CaptchaBase64 string `json:"captchaBase64"`
}

func getCaptchaApi(ctx server.Context) {
	var resp app.ResponseWrapper
	var err error

	id, s, answer, err := viewToken.GenCaptcha()
	if err != nil {
		log.Info("get captcha error")
		resp.Code = http.StatusInternalServerError
		ctx.Json(resp)
		return
	}

	log.Info("captcha answer: %s", answer)
	err = storage.Save(captchaPrefix+id, answer, time.Minute*10)
	if err != nil {
		log.Info("save captcha error")
		resp.Code = http.StatusInternalServerError
		ctx.Json(resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Data = respCaptcha{
		CaptchaBase64: s,
		CaptchaKey:    id,
	}
	ctx.Json(resp)
}
