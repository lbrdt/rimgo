package pages

import (
	"codeberg.org/video-prize-ranch/rimgo/api"
	"github.com/gofiber/fiber/v2"
)

func HandleGallery(c *fiber.Ctx) error {
	c.Set("Cache-Control", "public,max-age=604800")
	c.Set("X-Frame-Options", "DENY")
	c.Set("Referrer-Policy", "no-referrer")
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Robots-Tag", "noindex, noimageindex, nofollow")
	c.Set("Strict-Transport-Security", "max-age=31557600")
	c.Set("Permissions-Policy", "accelerometer=(), ambient-light-sensor=(), autoplay=(), battery=(), camera=(), cross-origin-isolated=(), display-capture=(), document-domain=(), encrypted-media=(), execution-while-not-rendered=(), execution-while-out-of-viewport=(), fullscreen=(), geolocation=(), gyroscope=(), interest-cohort=(), magnetometer=(), microphone=(), midi=(), navigation-override=(), payment=(), picture-in-picture=(), publickey-credentials-get=(), screen-wake-lock=(), sync-xhr=(), usb=(), web-share=(), xr-spatial-tracking=()")
	c.Set("Content-Security-Policy", "default-src 'none'; media-src 'self'; style-src 'self'; script-src 'none'; img-src 'self'; font-src 'self'; block-all-mixed-content; manifest-src 'self'")

	album, err := api.FetchAlbum(c.Params("galleryID"))
	if err != nil {
		return err
	}

	return c.Render("gallery", fiber.Map{
		"album": album,
		"isAlbum": false,
	})
}
