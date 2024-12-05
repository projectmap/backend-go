package product

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/utils"

	"github.com/gin-gonic/gin"
)

// ProductController data type
type Controller struct {
	service *Service
	logger  framework.Logger
	env     *framework.Env
}

type URLObject struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// New product Controller creates new product controller
func NewController(
	productService *Service,
	logger framework.Logger,
	env *framework.Env,
) *Controller {
	return &Controller{
		service: productService,
		logger:  logger,
		env:     env,
	}
}

//get all product

func (u *Controller) GetAllProduct(c *gin.Context) {
	product, err := u.service.GetAllProduct()
	if err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{
		"data": product,
	})

}

//add product

func (u *Controller) AddProduct(c *gin.Context) {
	var order models.Product

	if err := c.Bind(&order); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	// check if the product already exists

	if err := u.service.CreateProduct(&order); err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "product created"})
}
