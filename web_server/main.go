package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func main() {
	app := fiber.New()

	// use compress middleware
	app.Use(compress.New())

	app.Static("/", "../public")

	log.Fatal(app.Listen(":3000"))
}
