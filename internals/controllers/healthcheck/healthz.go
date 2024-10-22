package healthcheck

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (c *Controller) Healthz(ctx *fiber.Ctx) error {
	if !c.gitlab.Healthcheck() {
		return ctx.Status(http.StatusBadGateway).JSON("Not Ready")
	}
	return ctx.Status(http.StatusOK).JSON("OK")
}
