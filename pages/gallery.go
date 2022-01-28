package pages

import (
	"codeberg.org/video-prize-ranch/rimgo/api"
	"codeberg.org/video-prize-ranch/rimgo/types"
	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleGallery(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("Content-Security-Policy", "default-src 'none'; style-src 'self' *.cloudfront.net; media-src 'self' *.cloudfront.net; img-src 'self' *.cloudfront.net; font-src 'self' *.cloudfront.net; block-all-mixed-content")

	album, err := api.FetchAlbum(c.Params("galleryID"))
	if err != nil {
		return err
	}

	comments := []types.Comment{}
	if album.SharedWithCommunity {
		c.Set("Cache-Control", "public,max-age=604800")
		comments, err = api.FetchComments(c.Params("galleryID"))
		if err != nil {
			return err
		}
	} else {
		c.Set("Cache-Control", "public,max-age=31557600")
	}

	return c.Render("gallery", fiber.Map{
		"album":    album,
		"comments": comments,
	})
}
