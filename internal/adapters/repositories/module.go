package repositories

import (
	"github.com/volvofixthis/pow-server/internal/adapters/repositories/passage"
	"github.com/volvofixthis/pow-server/internal/adapters/repositories/pow"
	"go.uber.org/fx"
)

var Module = fx.Options(
	passage.Module,
	pow.Module,
)
