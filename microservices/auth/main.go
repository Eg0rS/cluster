package main

import (
	"auth/config"
	"auth/dal"
	"auth/utils/dbUtils"
	"context"
	"log"
)

func main() {
	log.SetFlags(log.Flags() | log.Llongfile)

	settings := config.Read()

	mainCtx := context.Background()

	done := make(chan struct{})
	db, err := dbUtils.NewPostgres(settings.DbConnectionString)
	if err != nil {
		panic(err)
	}

	server := serviceProvider{
		settings:               settings,
		done:                   done,
		userRepository:         dal.NewDbUserRepository(db),
		refreshTokenRepository: dal.NewDbRefreshTokenRepository(db),
		mainCtx:                mainCtx,
	}.provide()
	server.Start()
	<-done
}
