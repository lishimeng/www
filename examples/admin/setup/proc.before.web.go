package setup

import (
	"context"
	"github.com/lishimeng/www/examples/internal/domain"
)

func ComponentBeforeWeb(_ context.Context) error {
	domain.InitManagers()
	return nil
}
