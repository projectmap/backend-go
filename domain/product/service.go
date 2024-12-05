package product

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
)

// product Service service layer
type Service struct {
	logger     framework.Logger
	repository Repository
}

// New product Service creates a new product service
func NewService(
	logger framework.Logger,
	productRepository Repository,
) *Service {
	return &Service{
		logger:     logger,
		repository: productRepository,
	}
}

//get all order

func (s Service) GetAllProduct() (order []ProductSerializer, err error) {
	return s.repository.GetAllProduct()
}

//create product

func (s Service) CreateProduct(product *models.Product) error {
	return s.repository.Create(product).Error
}
