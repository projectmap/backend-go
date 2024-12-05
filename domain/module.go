package domain

import (
	"clean-architecture/domain/product"
	"clean-architecture/domain/user"

	"go.uber.org/fx"
)

var Module = fx.Module("domain",
	fx.Options(
		user.Module,
		product.Module,
	),
)
