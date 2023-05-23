package api

import (
	"auth/config"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
)

func NewAuthServer(
	settings *config.Settings,
	stopped chan struct{},
	mux http.Handler,
) *Server {
	return &Server{
		stopped:  stopped,
		settings: settings,
		mux:      mux,
	}
}

type Server struct {
	stopped    chan struct{}
	settings   *config.Settings
	mux        http.Handler
	httpServer *http.Server
}

func (s *Server) Start() {
	host := fmt.Sprintf(":%s", strconv.Itoa(s.settings.Port))
	tcpListener, err := net.Listen("tcp", host)
	if err != nil {
		log.Panicf("TCP listener wasn't created: %s", err)
	}

	s.httpServer = &http.Server{
		Addr:    host,
		Handler: s.mux,
	}
	go s.httpServer.Serve(tcpListener)
	fmt.Printf("HTTP server started on %d\n", s.settings.Port)
}

func (s *Server) Stop() {
	s.httpServer.Close()

	fmt.Printf("HTTP server stopped\n")
	if s.stopped != nil {
		s.stopped <- struct{}{}
	}
}
