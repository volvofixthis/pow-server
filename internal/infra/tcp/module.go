package tcp

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"TCP server",
	fx.Provide(NewTCPServer),
	fx.Invoke(
		registerLifecycleHooks,
	),
)

func registerLifecycleHooks(lc fx.Lifecycle, s *TCPServer) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go StartServer(s)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return StopServer(s)
		},
	})
}
