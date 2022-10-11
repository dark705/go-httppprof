package httppprof

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof" //nolint:gosec
	"time"
)

type server struct {
	httpServer *http.Server
	config     config
}

type config struct {
	httpListen string
}

func NewServer(httpListen string) *server { //nolint:golint,revive
	return &server{
		config: config{
			httpListen: httpListen,
		},
		httpServer: &http.Server{},
	}
}

func (s *server) Run() {
	log.Printf("httpPprof, start on: %s", s.config.httpListen)
	listener, err := net.Listen("tcp", s.config.httpListen)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("%s: %s", "httpPprof, fail open port", err)
	}

	go func() {
		err = s.httpServer.Serve(listener)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("%s: %s", "httpPprof, fail start", err)
		}
	}()
}

func (s *server) Shutdown() {
	log.Println("httpPprof, shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) //nolint:gomnd
	defer cancel()
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Printf("%s: %s", "httpPprof, fail shutdown", err)

		return
	}
	log.Println("httpPprof, success shutdown")
}
