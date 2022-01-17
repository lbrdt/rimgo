package pages

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HandleMedia(c *fiber.Ctx) error {
	res, err := FetchMedia(c.Params("baseName") + "." + c.Params("extension"))
	if err != nil {
		return err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	c.Set("Content-Type", res.Header.Get("Content-Type"));
	c.Write(data)
	return nil
}

func FetchMedia(filename string) (*http.Response, error) {
	res, err := http.Get("https://i.imgur.com/" + filename)
	return res, err
}