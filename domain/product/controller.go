package product

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/utils"
	"errors"

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

// get product by id
func (u *Controller) GetProductByID(c *gin.Context) {
	productID := c.Param("id")
	if productID == "" {
		utils.HandleValidationError(u.logger, c, ErrInvalidProductID)
		return
	}

	user, err := u.service.GetProductByID(productID)
	if err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

//update product

func (u *Controller) UpdateProduct(c *gin.Context) {

	productID := c.Param("id")
	if productID == "" {
		utils.HandleValidationError(u.logger, c, errors.New("ID is required"))
		return
	}

	var product ProductSerializer

	if err := c.ShouldBindJSON(&product); err != nil {
		utils.HandleValidationError(u.logger, c, err)
		return
	}
	if product.ProductName == "" && product.Price == 0 {
		utils.HandleValidationError(u.logger, c, errors.New("update data is required"))
		return
	}

	err := u.service.UpdateProduct(productID, product)
	if err != nil {
		utils.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "product updated"})
}
