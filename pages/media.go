package pages

import (
	"io/ioutil"
	"net/http"
	
	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleMedia(c *fiber.Ctx) error {
	c.Set("Cache-Control", "public,max-age=31557600")
	return handleMedia(c, "https://i.imgur.com/" + c.Params("baseName") + "." + c.Params("extension"))
}

func HandleUserCover(c *fiber.Ctx) error {
	c.Set("Cache-Control", "public,max-age=604800")
	return handleMedia(c, "https://imgur.com/user/" + c.Params("userID")  + "/cover?maxwidth=2560")
};

func HandleUserAvatar(c *fiber.Ctx) error {
	c.Set("Cache-Control", "public,max-age=604800")
	return handleMedia(c, "https://imgur.com/user/" + c.Params("userID")  + "/avatar")
};

func handleMedia(c *fiber.Ctx, url string) error {
	utils.SetHeaders(c)
	c.Set("Content-Security-Policy", "default-src 'none'; media-src 'self'; style-src 'self'; img-src 'self'; font-src 'self'; block-all-mixed-content")

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	c.Set("Content-Type", res.Header.Get("Content-Type"));
	return c.Send(body)
}