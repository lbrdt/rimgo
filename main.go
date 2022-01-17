package main

import (
	"log"
	"net/http"

	"codeberg.org/video-prize-ranch/go-rimgu/pages"
	"codeberg.org/video-prize-ranch/go-rimgu/static"
	"codeberg.org/video-prize-ranch/go-rimgu/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/handlebars"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("/etc/rimgu/")
	viper.AddConfigPath("$HOME/.config/rimgu")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("RIMGU_PORT", "3000")
	viper.SetDefault("RIMGU_HOST", "localhost")
	viper.SetDefault("RIMGU_ADDRESS", "0.0.0.0")
	viper.SetDefault("RIMGU_IMGUR_CLIENT_ID", "546c25a59c58ad7")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	
	engine := handlebars.NewFileSystem(http.FS(views.GetFiles()), ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
		Prefork: viper.GetBool("FIBER_PREFORK"),
		UnescapePath: true,
		StreamRequestBody: true,
	})

	app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.FS(static.GetFiles()),
	}))

	app.Get("/media/:baseName.:extension", pages.HandleMedia)
	app.Get("/a/:albumID", pages.HandleAlbum)
	app.Get("/t/:tagID", pages.HandleAlbum)
	/*app.Get("/user/:userID", pages.HandleUser)
	app.Get("/user/:userID/cover", pages.HandleUserCover)
	app.Get("/gallery/:galleryID", pages.HandleGallery)*/

	app.Listen(":" + viper.GetString("RIMGU_PORT"))
}