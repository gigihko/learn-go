package src

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	// Render index template
	return c.Render("index", fiber.Map{
		"Title":       "TEST HTML",
		"Description": "IKHSAN ACADEMY",
		"Greeting":    "Hello, world!",
	})
}

func Dynamic(c *fiber.Ctx) error {
	name := "guest"
	if c.Params("name") != "" {
		name = c.Params("name")
	}
	return c.SendString("Hello " + name)
}
