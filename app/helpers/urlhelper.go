package helper

import (
	"strconv"

	"github.com/RiddhiNahata/shorten-url/app/utils"
	"github.com/gofiber/fiber/v2"
)

type requestShortURlBody struct {
	LargeURL string `json:"url"`
	Email    string `json:"email"`
}

// redirect to orginal url after fetching from DB
func (h handler) RedirectShortURL(c *fiber.Ctx) error {

	existingUrl, err := utils.GetLargeUrlFromShort(h.DB, c.BaseURL()+c.OriginalURL())
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = utils.UpdateURLCounter(h.DB, existingUrl)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Redirect(existingUrl.LargeUrl)
}

// Adding a user to DB and creating the short url
func (h handler) GetShortenURL(c *fiber.Ctx) error {

	body := requestShortURlBody{}

	// parse body, attach to requestShortURlBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userID, err := utils.CreateUser(h.DB, body.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	shortUrl, err := utils.CreateNewShortUrl(h.DB, body.LargeURL, userID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&shortUrl)
}

// Get all urls by User id
func (h handler) GetUrlsByUser(c *fiber.Ctx) error {

	id := c.Query("userid")
	userid, err := strconv.Atoi(id)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	urls, err := utils.GetUrlsByUserID(h.DB, userid)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&urls)
}

// Delete the urls by primary key
func (h handler) DeleteUrls(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = utils.DeleteUrls(h.DB, id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Deleted")
}

// Find the domain hi count on descending order eg google 5, facebook 5, twitter 2 etc
func (h handler) GetDomainCount(c *fiber.Ctx) error {

	id := c.Query("userid")
	userid, err := strconv.Atoi(id)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	urls, err := utils.GetDomainCountByUserID(h.DB, userid)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&urls)
}
