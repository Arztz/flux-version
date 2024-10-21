package http

import "github.com/gofiber/fiber/v2"

func (s *Server) configure() {
	s.fiber = fiber.New()
	s.fiber.Get("/version", s.controller.gitlab.GetJson)
	s.fiber.Get("/healthz", s.controller.healthcheck.Healthz)
}
