package product

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
)

// UserRepository database structure
type Repository struct {
	infrastructure.Database
	logger framework.Logger
}

// NewUserRepository creates a new product repository
func NewRepository(db infrastructure.Database, logger framework.Logger) Repository {
	return Repository{db, logger}
}

func Migrate(r Repository) error {
	r.logger.Info("[Migrating...products]")

	if err := r.DB.AutoMigrate(&models.Product{}); err != nil {
		r.logger.Error("[Migration failed...product]")
		return err
	}
	return nil
}

// get all order
func (r *Repository) GetAllProduct() (order []ProductSerializer, err error) {
	query := r.Model(&models.Product{})
	return order, query.Find(&order).Error

}

//get product by id

func (r *Repository) GetProductByID(productID string) (product ProductSerializer, err error) {
	query := r.Model(&models.Product{})
	return product, query.First(&product, "id = ?", productID).Error

}

//update product

func (r *Repository) UpdateProduct(productID string, product ProductSerializer) (err error) {

	query := r.Model(&models.Product{})
	productData := map[string]interface{}{}

	// Conditionally add fields to the update data if they are not empty
	if product.ProductName != "" {
		productData["product_name"] = product.ProductName
	}
	if product.Quantity != 0 {
		productData["quantity"] = product.Quantity
	}
	if product.Price != 0 {
		productData["price"] = product.Price
	}

	return query.Where("id = ?", productID).Updates(productData).Error
}
