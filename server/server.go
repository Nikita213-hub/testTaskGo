package server

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	address string
	port    string
	server  *http.Server
	mux     *http.ServeMux
}

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Router interface {
	Routes() []Route
}

func NewServer(address, port string) *Server {
	return &Server{
		address: address,
		port:    port,
		mux:     http.NewServeMux(),
	}
}

func (s *Server) Run(ctx context.Context, routers ...Router) error {
	for _, router := range routers {
		for _, route := range router.Routes() {
			s.mux.HandleFunc(route.Method+" "+route.Path, route.Handler)
		}
	}

	s.server = &http.Server{
		Addr:    s.address + s.port,
		Handler: s.mux,
	}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	fmt.Printf("Server is listening %v\n", s.address+s.port)
	<-ctx.Done()

	return s.server.Shutdown(context.Background())
}
