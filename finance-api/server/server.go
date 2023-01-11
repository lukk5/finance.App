package main

import (
	"context"
	"finance-api-v1/controllers"
	"finance-api-v1/core/config"
	"finance-api-v1/core/database"
	userRepo "finance-api-v1/core/database/repo"
	"finance-api-v1/core/middleware/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"log"
	"net/http"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		runApplication()
	},
}

func newServer(lc fx.Lifecycle) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", "8080"),
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("Start to rest api server : 8080")
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Printf("Stopped rest api server")
			return srv.Shutdown(ctx)
		},
	})
	return r
}

func printAppInfo(cfg *config.Config) {
	log.Printf("application information", "config", cfg)
}

func loadConfig() (*config.Config, error) {
	return config.Load(configFile)
}

func runApplication() {
	// setup (DI) and run
	app := fx.New(
		fx.Provide(
			// setup config
			loadConfig,
			// setup database
			database.NewDatabase,
			// setup repo
			userRepo.NewUserRepo,
			// setup handlers
			handlers.NewUserHandler,
			// server
			newServer,
		),
		fx.Invoke(
			printAppInfo,
			controllers.UserController,
		),
	)
	app.Run()
}
