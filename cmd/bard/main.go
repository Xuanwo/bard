package main

import (
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Xuanwo/bard/contexts"
	"github.com/Xuanwo/bard/handler"
)

// Cmd is the main command for bard.
var Cmd = cobra.Command{
	Use: "bard",
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}

		viper.SetConfigName("config")
		viper.AddConfigPath(configPath)
		err = viper.ReadInConfig()
		if err != nil {
			return err
		}

		err = contexts.Setup()
		if err != nil {
			return err
		}

		app := iris.New()
		app.Configure(
			iris.WithFireMethodNotAllowed,
			iris.WithPostMaxMemory(128<<20), // Max file is 128MB
		)
		app.Use(
			recover2.New(),
			logger.New(),
		)

		app.Get("/ping", func(ctx iris.Context) {
			_, _ = ctx.JSON(iris.Map{
				"message": "pong",
			})
		})
		app.Post("/", handler.Create)
		app.Get("/{short_id:string max(6)}", handler.Get)

		err = app.Run(iris.Addr(contexts.Server.Listen))
		if err != nil {
			log.Fatalf("Server existed: %v", err)
		}
		return nil
	},
}

func init() {
	Cmd.Flags().String("config", "", "")
}

func main() {
	err := Cmd.Execute()
	if err != nil {
		log.Fatalf("bard exited: %v", err)
	}
}
