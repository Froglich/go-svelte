package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type pageData struct {
	Username string `json:"username"`
}

func indexHandler(c *fiber.Ctx) error {
	return c.Render("index", pageData{
		Username: "hello world",
	})
}

func main() {
	templateEngine := html.New("./static", ".html")
	templateEngine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: templateEngine,
	})

	app.Get("/", indexHandler)
	app.Static("/_app", "./static/_app")
	app.Static("/global", "./static/global")

	log.Fatal(app.Listen("127.0.0.1:8000"))
}
