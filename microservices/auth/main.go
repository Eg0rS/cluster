package main

import (
	"auth/config"
	"auth/dal"
	"auth/utils/dbUtils"
	"context"
	"log"

	_ "auth/docs"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

//	@title			Swagger of API
//	@version		1.0
//	@description	This is a sample server celler server.
//
// @BasePath	/
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
