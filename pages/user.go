package pages

import (
	"strconv"
	"sync"

	"codeberg.org/video-prize-ranch/rimgo/api"
	"codeberg.org/video-prize-ranch/rimgo/types"
	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleUser(c *fiber.Ctx) error {
	utils.SetHeaders(c)
	c.Set("Cache-Control", "public,max-age=604800")
	c.Set("Content-Security-Policy", "default-src 'none'; media-src 'self'; style-src 'unsafe-inline' 'self'; img-src 'self'; font-src 'self'; block-all-mixed-content")

	page := "0"
	if c.Query("page") != "" {
		page = c.Query("page")
	}

	pageNumber, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		pageNumber = 0
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	user, err := types.User{}, error(nil)
	go func() {
		defer wg.Done()
		user, err = api.FetchUser(c.Params("userID"))
	}()
	if err != nil {
		return err
	}

	submissions, err := []types.Submission{}, error(nil)
	go func() {
		defer wg.Done()
		submissions, err = api.FetchSubmissions(c.Params("userID"), "newest", page)
	}()
	if err != nil {
		return err
	}

	wg.Wait()
	return c.Render("user", fiber.Map{
		"user":        user,
		"submissions": submissions,
		"page":        page,
		"nextPage": 	 pageNumber + 1,
		"prevPage": 	 pageNumber - 1,
	})
}
