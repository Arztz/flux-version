package gitlab

import (
	"flux-version/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (c *Controller) GetJson(ctx *fiber.Ctx) error {
	var all []types.Project
	for _, p := range c.config.ProjectList {
		category := c.service.ReadFile(p)
		project := types.Project{Project: p, Category: []types.Category{}}
		res, err := c.service.GenerateJSON(project, category)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(map[string]any{"error": err.Error()})
		}
		if all == nil {
			all = []types.Project{}
		}

		all = append(all, res)
	}
	ctx.Response().Header.Set("Content-Type", "application/json; charset=utf-8")
	return ctx.Status(http.StatusOK).JSON(all)

}
