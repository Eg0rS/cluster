package api

import (
	"auth-gateway/config"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

func NewServer(
	settings *config.Settings,
	mux *http.ServeMux,
	stopped chan struct{},
) *Server {
	return &Server{
		settings: settings,
		mux:      mux,
		stopped:  stopped,
	}
}

type Server struct {
	settings   *config.Settings
	mux        *http.ServeMux
	httpServer *http.Server
	stopped    chan struct{}
}

func (s *Server) Start() {
	host := fmt.Sprintf(":%s", strconv.Itoa(s.settings.Port))
	tcpListener, err := net.Listen("tcp", host)
	if err != nil {
		panic(fmt.Errorf(
			"TCP listener wasn't created: %s",
			err.Error(),
		))
	}

	s.httpServer = &http.Server{
		Addr:         host,
		Handler:      s.mux,
		ReadTimeout:  time.Duration(s.settings.ServerTimeout) * time.Second,
		WriteTimeout: time.Duration(s.settings.RequestForwarderTimeout) * time.Second,
	}
	s.httpServer.SetKeepAlivesEnabled(false)
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
