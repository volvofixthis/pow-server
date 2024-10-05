package core

import (
	"github.com/volvofixthis/pow-server/internal/core/ports"
	"github.com/volvofixthis/pow-server/internal/core/services/passage"
	"github.com/volvofixthis/pow-server/internal/core/services/pow"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(passage.NewPassageService, fx.As(new(ports.PassageService))),
	),
	pow.Module,
)
