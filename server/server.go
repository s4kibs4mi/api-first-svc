package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/s4kibs4mi/api-first-svc/configs"
)

type Server interface {
	Register()
	Start() error
	Stop(ctx context.Context) error
}

func NewFiberServer(cfg *configs.Config, fiberApp *fiber.App) Server {
	return &serverImpl{
		cfg:      cfg,
		fiberApp: fiberApp,
	}
}

type serverImpl struct {
	cfg      *configs.Config
	fiberApp *fiber.App
}

const (
	scalarDocsLoaderScript = `
<!doctype html>
<html>
  <head>
    <title>API Reference</title>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <script
      id="api-reference"
      data-url="/openapi.json"></script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
  </body>
</html>`
)

func (s *serverImpl) Register() {
	s.fiberApp.Get("/playground", func(ctx *fiber.Ctx) error {
		ctx.Set("Content-Type", "text/html")
		ctx.Write([]byte(scalarDocsLoaderScript))
		return nil
	})
}

func (s *serverImpl) Start() error {
	return s.fiberApp.Listen(fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port))
}

func (s *serverImpl) Stop(ctx context.Context) error {
	return s.fiberApp.ShutdownWithContext(ctx)
}
