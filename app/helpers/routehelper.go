package helper

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	// declare all api routes
	routes := app.Group("/urls")
	routes.Post("/shorten", h.GetShortenURL)
	routes.Get("/list", h.GetUrlsByUser)
	routes.Get("/redirect", h.RedirectShortURL)
	routes.Delete("/delete", h.DeleteUrls)
	routes.Get("/domain", h.GetDomainCount)

	// special route for short url redirect
	app.Get("/*", h.RedirectShortURL)

}
