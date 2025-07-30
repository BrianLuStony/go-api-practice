package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type ServerConfig struct {
	Addr   string
	Router http.Handler
}

type AppServer struct {
	server *http.Server
}

func NewServer(config ServerConfig) *AppServer {
	srv := &http.Server{
		Addr:         config.Addr,
		Handler:      config.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	return &AppServer{server: srv}
}

func (s *AppServer) Start() error {
	log.Println("Starting server on ", s.server.Addr)
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
			return
		}
	}()
	return nil
}
func (s *AppServer) ShutdownGracefully(timeout time.Duration) error {
	log.Println("Shutting down server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed: %v", err)
		return err
	}
	log.Println("Server stopped gracefully")
	return nil
}
