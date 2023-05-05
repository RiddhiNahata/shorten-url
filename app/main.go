package main

import (
	"log"

	"github.com/RiddhiNahata/shorten-url/app/common/config"
	"github.com/RiddhiNahata/shorten-url/app/common/db"
	helper "github.com/RiddhiNahata/shorten-url/app/helpers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	h := db.Init(&c)
	app := fiber.New()

	// Register routes
	helper.RegisterRoutes(app, h)

	log.Fatal(app.Listen(":3000"))
}
