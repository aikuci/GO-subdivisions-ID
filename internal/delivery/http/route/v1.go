package route

import "github.com/gofiber/fiber/v2"

func (c *RouteConfig) SetupV1Route() {
	v1 := c.App.Group("/v1")

	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG from V1")
	})

	v1.Get("/provinces", c.ProvinceController.List)
	v1.Get("/provinces/:id<int>", c.ProvinceController.GetByID)
}
