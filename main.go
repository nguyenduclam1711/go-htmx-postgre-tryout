package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/nguyenduclam1711/go-htmx-postgre-tryout/database"
	"github.com/nguyenduclam1711/go-htmx-postgre-tryout/env"
)

func main() {
	dbCtx := context.Background()

	// load env
	env.LoadEnv()

	// connect to db
	database.ConnectDB(dbCtx)

	// generates db tables
	database.GenerateTables()

	db := database.Db
	defer db.Close(dbCtx)

	app := fiber.New()

	// use compress middleware
	app.Use(compress.New())

	app.Static("/", "../public")

	log.Fatal(app.Listen(":3000"))
}
