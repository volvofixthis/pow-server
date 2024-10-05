package http

import (
	"context"

	"github.com/labstack/echo/v4"
	httpadapter "github.com/volvofixthis/pow-server/internal/adapters/handlers/http"
	"github.com/volvofixthis/pow-server/internal/infra/config"
	"go.uber.org/zap"
)

type HTTPServer struct {
	e   *echo.Echo
	c   *config.AppConfig
	log *zap.Logger
}

func NewHTTPServer(
	log *zap.Logger,
	config *config.AppConfig,
	powTaskAdapter *httpadapter.PowTaskAdapter,
	passageAdapter *httpadapter.PassageAdapter,
) *HTTPServer {
	e := echo.New()
	server := HTTPServer{e: e, c: config, log: log}

	e.POST("/v1/pow", powTaskAdapter.CreateTask)
	e.POST("/v1/passage", passageAdapter.GetPassage)

	return &server
}

func StartServer(server *HTTPServer) error {
	go func() {
		if err := server.e.Start(server.c.ApiAddress); err != nil {
			server.log.Error("Failed to start HTTP server", zap.Error(err))
		}
	}()
	return nil
}

func StopServer(ctx context.Context, server *HTTPServer) error {
	return server.e.Shutdown(ctx)
}
