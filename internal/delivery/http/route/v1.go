package route

import "github.com/gofiber/fiber/v2"

func (c *RouteConfig) SetupV1Route() {
	v1 := c.App.Group("/v1")

	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG from V1")
	})

	v1.Get("/provinces", c.ProvinceController.List)
	v1.Get("/provinces/:id<int>", c.ProvinceController.GetFirstByID)

	v1.Get("/provinces/:id_province<int>/cities", c.CityController.GetByIdProvince)
	v1.Get("/provinces/:id_province<int>/cities/:id<int>", c.CityController.GetByIDAndIdProvince)
	v1.Get("/cities", c.CityController.CrudController.List)
	v1.Get("/cities/:id<int>", c.CityController.CrudController.GetByID)
}
