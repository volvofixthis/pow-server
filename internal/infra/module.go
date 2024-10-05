package infra

import (
	"github.com/volvofixthis/pow-server/internal/infra/logging"
	"github.com/volvofixthis/pow-server/internal/infra/tcp"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"infra",
	logging.Module,
	// http.Module,
	tcp.Module,
)
