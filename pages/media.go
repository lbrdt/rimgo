package pages

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func HandleMedia(c *fiber.Ctx) error {
	res, err := http.Get("https://i.imgur.com/" + c.Params("baseName") + "." + c.Params("extension"))
	if err != nil {
		return err
	}

	c.Set("Content-Type", res.Header.Get("Content-Type"));
	contentLen, _ := strconv.Atoi(res.Header.Get("Content-Length"))
	c.SendStream(res.Body, contentLen)
	return nil
}