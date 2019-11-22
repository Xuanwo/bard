package main

import (
	"log"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/ping", func(ctx iris.Context) {
		_, _ = ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		log.Fatalf("Server existed: %v", err)
	}
}
