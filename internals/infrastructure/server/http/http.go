package http

import (
	"fmt"
	log "log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (s *Server) Start() {
	s.configure()

	go func() {
		if err := s.fiber.Listen(fmt.Sprintf(":%s", s.config.HttpPort)); err != nil {
			log.Error("Error, HTTP server failed to listen and serve:", err)
			panic(err)
		}
	}()

	log.Info("HTTP server listening and serving on port: %s", s.config.HttpPort)

	gracefulStop := make(chan os.Signal, 3)
	signal.Notify(gracefulStop, os.Interrupt)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGTERM)

	<-gracefulStop
	close(gracefulStop)

	s.Shutdown()
}

func (s *Server) Shutdown() {
	timeout := time.Duration(s.config.HTTPServerTimeout) * time.Second

	log.Info("Shutting down HTTP server")
	if err := s.fiber.ShutdownWithTimeout(timeout); err != nil {
		log.Error("HTTP server shutdown: %v", err)
		return
	}

	log.Info(fmt.Sprintf("HTTP Server Timeout of %s", timeout.String()))
	log.Info("HTTP server gracefully stopped")
}
