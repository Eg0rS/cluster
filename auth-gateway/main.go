package main

import (
	"auth-gateway/config"
	"log"
)

func main() {
	log.SetFlags(log.Flags() | log.Llongfile)
	settings := config.Read()

	done := make(chan struct{})
	server := serviceProvider{
		settings,
		done,
	}.provideServer()
	server.Start()
	<-done
}
