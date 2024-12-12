package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController data type
type Controller struct {
	service *Service
	logger  framework.Logger
	env     *framework.Env
}

type URLObject struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// NewUserController creates new user controller
func NewController(
	userService *Service,
	logger framework.Logger,
	env *framework.Env,
) *Controller {
	return &Controller{
		service: userService,
		logger:  logger,
		env:     env,
	}
}

// CreateUser creates the new user
func (u *Controller) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	// check if the user already exists

	if err := u.service.Create(&user); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}

//create order

func (u *Controller) CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.Bind(&order); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	// check if the order already exists

	if err := u.service.CreateOrder(&order); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "order created"})
}

// GetOneUser gets one user
func (u *Controller) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		utils.HandleValidationError(u.logger, c, ErrInvalidUserID)
		return
	}

	user, err := u.service.GetUserByID(userID)
	if err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

//delete order

func (u *Controller) DeleteOrderByID(c *gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		utils.HandleValidationError(u.logger, c, ErrInvalidOrderID)
		return
	}

	err := u.service.DeleteOrderByID(orderId)
	if err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "order deleted"})

}

//get all order

func (u *Controller) GetAllOrder(c *gin.Context) {

	order, err := u.service.GetAllOrder()
	if err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{
		"data": order,
	})

}

//get order by product

func (u *Controller) GetTotalOrderForProduct(c *gin.Context) {
	quantityAbove := c.Query("quantity_above")

	var (
		quantityAboveInt int
	)

	quantityAboveInt, _ = strconv.Atoi(quantityAbove)

	if quantityAbove == "" {
		quantityAboveInt = 0
	}

	filter := OrderGroupListFilter{

		QuantityAbove: quantityAboveInt,
	}
	order, err := u.service.GetTotalOrderForProduct(filter)
	if err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{
		"data": order,
	})

}
