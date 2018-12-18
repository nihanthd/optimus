package pins

import (
	"github.com/nihanthd/optimus/pins/handler"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	handler.NewHandler,
	handler.NewPinsHandler,
)
