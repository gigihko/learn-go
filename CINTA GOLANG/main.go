package main

import (
	"cinta_golang/src"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./static")

	app.Get("/", src.Hello)
	app.Get("/nama/:name?", src.Dynamic)

	log.Fatal(app.Listen(":3000"))
}
