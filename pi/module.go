package pi

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewPI,
)
