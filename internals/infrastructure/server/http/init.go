package http

import (
	"flux-version/internals/config"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	controller *Controller
	config     config.Configuration
	fiber      *fiber.App
}

func NewServer(controller *Controller, config config.Configuration) *Server {
	return &Server{
		controller: controller,
		config:     config,
	}
}
