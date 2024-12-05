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
