package pages

import (
	"codeberg.org/video-prize-ranch/go-rimgu/api"
	"github.com/gofiber/fiber/v2"

)

func HandleAlbum(c *fiber.Ctx) error {
	// https://imgur.com/a/DfEsrAB

	album, err := api.FetchAlbum(c.Params("albumID"))
	if err != nil {
		return err
	}

	return c.Render("gallery", fiber.Map{
		"album": album,
		"isAlbum": true,
	})
}
