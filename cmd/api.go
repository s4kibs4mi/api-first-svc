package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/gofiber/fiber/v2"
	logger2 "github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/s4kibs4mi/api-first-svc/configs"
	"github.com/s4kibs4mi/api-first-svc/handlers"
	"github.com/s4kibs4mi/api-first-svc/handlers/user"
	"github.com/s4kibs4mi/api-first-svc/log"
	"github.com/s4kibs4mi/api-first-svc/server"
	userUC "github.com/s4kibs4mi/api-first-svc/usecases/user"
	"github.com/spf13/cobra"
)

var ApiCmd = &cobra.Command{
	Use: "api",
	Run: humacli.WithOptions(executeApi),
}

func executeApi(cmd *cobra.Command, args []string, options *configs.Config) {
	logger := log.New()

	logger.Info("Starting API First Service")

	fiberApp := fiber.New()
	fiberApp.Use(recover2.New(recover2.Config{
		EnableStackTrace: true,
	}))
	fiberApp.Use(logger2.New())

	humaCfg := huma.DefaultConfig("API First Service", "1.0.0")
	humaApi := humafiber.New(fiberApp, humaCfg)

	userUsecase := userUC.NewUseCase()

	var apiHandlers []handlers.Handler
	apiHandlers = append(apiHandlers, user.NewHandler(humaApi, userUsecase))
	for _, h := range apiHandlers {
		h.Register()
	}

	apiServer := server.NewFiberServer(options, fiberApp)
	apiServer.Register()

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGKILL, os.Interrupt)

	go func() {
		if err := apiServer.Start(); err != nil {
			logger.Fatal("Failed to start API server: ", err)
		}
	}()

	<-stop

	logger.Info("Stopping API server...")

	cancelCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_ = apiServer.Stop(cancelCtx)

	logger.Info("API server stopped.")
}
