package handlers

import (
	"github.com/volvofixthis/pow-server/internal/adapters/handlers/conn"
	"github.com/volvofixthis/pow-server/internal/adapters/handlers/http"
	"go.uber.org/fx"
)

var Module = fx.Options(
	http.Module,
	conn.Module,
)
