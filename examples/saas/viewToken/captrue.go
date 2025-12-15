package viewToken

import (
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore
var cp *base64Captcha.Captcha

func init() {
	driver := &base64Captcha.DriverString{
		Height:          40, // 与界面上的尺寸匹配
		Width:           120,
		Length:          4,
		Source:          "1234567890",
		ShowLineOptions: base64Captcha.OptionShowSineLine,
		NoiseCount:      0,
	}
	cp = base64Captcha.NewCaptcha(driver, store)

}

func GenCaptcha() (id, b64s, answer string, err error) {
	id, b64s, answer, err = cp.Generate()
	return
}
