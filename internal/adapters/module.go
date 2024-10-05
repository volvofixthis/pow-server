package adapters

import (
	"github.com/volvofixthis/pow-server/internal/adapters/handlers"
	"github.com/volvofixthis/pow-server/internal/adapters/repositories"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	repositories.Module,
)
