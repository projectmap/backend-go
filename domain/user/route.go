package user

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
)

// UserRoutes struct
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

// Setup user routes
func RegisterRoute(r *Route) {
	r.logger.Info("Setting up routes")

	api := r.handler.Group("/api")

	api.POST("/user", r.controller.CreateUser)
	api.POST("/order", r.controller.CreateOrder)
	api.GET("/user/:id", r.controller.GetUserByID)
	api.DELETE("/delete-order/:id", r.controller.DeleteOrderByID)
	api.GET("/order", r.controller.GetAllOrder)
	api.GET("/order-by-product", r.controller.GetTotalOrderForProduct)

}
