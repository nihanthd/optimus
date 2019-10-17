package ports

import (
	"github.com/nihanthd/optimus/ports/handler"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	handler.NewHandler,
	handler.NewPortsHandler,
)
