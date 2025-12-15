package viewToken

import (
	"testing"
	"time"
)

func TestGenCapture(t *testing.T) {

	GenCaptcha()
	time.Sleep(time.Second)
}
