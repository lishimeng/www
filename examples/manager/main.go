package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www/examples/internal/db"
	"github.com/lishimeng/www/examples/internal/etc"
	"github.com/lishimeng/www/examples/manager/ddd"
	"github.com/lishimeng/www/examples/manager/setup"
	"time"
)
import _ "github.com/lib/pq"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond * 50)
}

func _main() (err error) {

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(
			&etc.Config,
			app.WithDefaultCallback("config"),
		)
		if err != nil {
			return err
		}

		var dbConfig persistence.BaseConfig
		c := persistence.PostgresConfig{
			UserName:  etc.Config.Db.User,
			Password:  etc.Config.Db.Password,
			Host:      etc.Config.Db.Host,
			Port:      etc.Config.Db.Port,
			DbName:    etc.Config.Db.Database,
			InitDb:    true,
			AliasName: "default",
			SSL:       etc.Config.Db.Ssl,
		}
		dbConfig = c.Build()

		builder.EnableDatabase(dbConfig,
			db.RegisterTables()...).
			PrintVersion().
			EnableDatabaseLog().
			SetWebLogLevel("debug").
			EnableWeb(etc.Config.Web.Listen, ddd.Route).
			ComponentBefore(setup.ComponentBeforeWeb)

		return err
	}, func(s string) {
		log.Info(s)
	})

	return
}
