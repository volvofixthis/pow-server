package http

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"HTTP server",
	fx.Provide(NewHTTPServer),
	fx.Invoke(
		registerLifecycleHooks,
	),
)

func registerLifecycleHooks(lc fx.Lifecycle, s *HTTPServer) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			return StartServer(s)
		},
		OnStop: func(ctx context.Context) error {
			return StopServer(ctx, s)
		},
	})
}
