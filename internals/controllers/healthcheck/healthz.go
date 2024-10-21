package healthcheck

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (c *Controller) Healthz(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON("OK")

}
