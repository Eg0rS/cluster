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

	batchSize := 100

	done := make(chan struct{})
	dbConnector := dbUtils.NewDBConnector(settings.DbConnectionString)
	if err := dbConnector.Open(); err != nil {
		panic(err)
	}
	db, err := dbConnector.GetDB()
	if err != nil {
		panic(err)
	}
	clickConnect, err := clickhouse.ConfigureClickHouseConn(settings.ClickHouseRepository)
	if err != nil {
		log.Fatalln("Нет подключения к ClickHouse")
	}
	clickRepository := clickhouse.NewDBRepository(clickConnect)

	server := serviceProvider{
		settings:               settings,
		done:                   done,
		userRepository:         dal.NewDbUserRepository(db),
		userRoleRepository:     dal.NewDbUserRoleRepository(db),
		refreshTokenRepository: dal.NewDbRefreshTokenRepository(settings),
		passwordHasher:         passwordHasher.New(settings),
		loggerSlack:            loggerSlack.New(settings),
		availableUsersRepo:     dal.NewDbAvailableUsers(settings),
		uuidRepository:         dal.NewDbUUIDRepository(settings),
		clickRepository:        clickRepository,
		loggerService:          logger.New(settings),
		mainCtx:                mainCtx,
		authLog:                authLog,
	}.provide()
	server.Start()
	<-done
	dbConnector.Close()
}
