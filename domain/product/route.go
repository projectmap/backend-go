package product

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
)

// Product Routes struct
type Route struct {
	logger     framework.Logger
	handler    infrastructure.Router
	controller *Controller
}

func NewRoute(
	logger framework.Logger,
	handler infrastructure.Router,
	controller *Controller,
) *Route {
	return &Route{
		handler:    handler,
		logger:     logger,
		controller: controller,
	}

}

// Setup product routes
func RegisterRoute(r *Route) {
	r.logger.Info("Setting up routes")

	api := r.handler.Group("/api")

	api.POST("/product", r.controller.AddProduct)
	api.GET("/product", r.controller.GetAllProduct)
	api.GET("/product/:id", r.controller.GetProductByID)
	api.PATCH("/product/:id", r.controller.UpdateProduct)
	api.DELETE("/product/:id", r.controller.DeleteProduct)
	api.GET("/filter-product", r.controller.GetFilteredProduct)

}
