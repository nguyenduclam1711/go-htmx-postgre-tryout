package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/nguyenduclam1711/go-htmx-postgre-tryout/database"
	"github.com/nguyenduclam1711/go-htmx-postgre-tryout/env"
	"github.com/nguyenduclam1711/go-htmx-postgre-tryout/models"
	"github.com/nguyenduclam1711/go-htmx-postgre-tryout/repositories"
)

func main() {
	// load env
	env.LoadEnv()

	// connect to db
	database.ConnectDB()

	// generates db tables
	database.CreateTables()

	defer database.Db.Close()

	// generates all repositories
	repositories.GenerateMapRepositories()

	repositories.MapRepositories.User.Create(models.User{
		Username: "hahah",
		FullName: "heheh",
	})

	app := fiber.New()

	// use compress middleware
	app.Use(compress.New())

	app.Static("/", "public")

	log.Fatal(app.Listen(":3000"))
}
