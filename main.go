package main

import (
	"context"
	"fmt"
	"net/http"

	"codeberg.org/video-prize-ranch/rimgo/pages"
	"codeberg.org/video-prize-ranch/rimgo/static"
	"codeberg.org/video-prize-ranch/rimgo/views"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadaptor "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/handlebars"
	"github.com/spf13/viper"
)

var fiberLambda *fiberadaptor.FiberLambda

func init() {
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
		fmt.Println(err)
	}

	engine := handlebars.NewFileSystem(http.FS(views.GetFiles()), ".hbs")
	app := fiber.New(fiber.Config{
		Views:             engine,
		Prefork:           viper.GetBool("FIBER_PREFORK"),
		UnescapePath:      true,
		StreamRequestBody: false,
	})

	app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.FS(static.GetFiles()),
	}))

	app.Get("/robots.txt", func(c *fiber.Ctx) error {
		file, _ := static.GetFiles().ReadFile("robots.txt")
		_, err := c.Write(file)
		return err
	})

	app.Get("/", pages.FrontpageHandler)
	app.Get("/:baseName.:extension", pages.HandleMedia)
	app.Get("/a/:galleryID", pages.HandleGallery)
	//app.Get("/t/:tagID", pages.HandleAlbum)
	app.Get("/user/:userID", pages.HandleUser)
	app.Get("/user/:userID/cover", pages.HandleUserCover)
	app.Get("/user/:userID/avatar", pages.HandleUserAvatar)
	app.Get("/gallery/:galleryID", pages.HandleGallery)

	fiberLambda = fiberadaptor.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}