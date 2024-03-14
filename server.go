package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	api := app.Group("/api")
	chall := api.Group("/chall")

	chall.Get("/", func(ctx *fiber.Ctx) error {

	})
	chall.Get("/:id", func(ctx *fiber.Ctx) error {
		flagSubmitted := ctx.Query("flag", "None")

	})

	app.Listen("")
}
