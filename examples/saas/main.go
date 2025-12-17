package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/www/examples/internal/db"
	"github.com/lishimeng/www/examples/internal/etc"
	"github.com/lishimeng/www/examples/saas/ddd"
	"github.com/lishimeng/www/examples/saas/setup"
	"github.com/lishimeng/x/container"
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

	// todo: 系统group写在这里

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

		if etc.Config.Token.Enable {
			issuer := etc.Config.Token.Issuer
			tokenKey := []byte(etc.Config.Token.Key)
			builder = builder.EnableTokenValidator(func(inject app.TokenValidatorInjectFunc) {
				provider := token.NewJwtProvider(token.WithIssuer(issuer),
					token.WithKey(tokenKey, tokenKey), // hs256的秘钥必须是[]byte
					token.WithAlg("HS256"),
					token.WithDefaultTTL(time.Duration(etc.Config.Token.Ttl)*time.Minute),
				)
				var storage token.Storage
				if etc.Config.Token.EnableJwtTokenCache {
					// 使用 Redis 缓存验证 jwt token
					storage = token.NewRedisStorage(app.GetCache())
				} else {
					storage = token.NewLocalStorage(provider)
				}
				container.Add(provider)
				inject(storage)
			})
		}

		builder.EnableDatabase(dbConfig,
			db.RegisterTables()...).
			PrintVersion().
			EnableDatabaseLog().
			SetWebLogLevel("debug").
			EnableWeb(etc.Config.Web.Listen, ddd.Route).
			//EnableStaticWeb(func() http.FileSystem {
			//	return http.FS(static.Static)
			//}).
			ComponentBefore(setup.ComponentBeforeWeb)

		return err
	}, func(s string) {
		log.Info(s)
	})

	return
}
