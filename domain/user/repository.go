package user

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

// NewUserRepository creates a new user repository
func NewRepository(db infrastructure.Database, logger framework.Logger) Repository {
	return Repository{db, logger}
}

func Migrate(r Repository) error {
	r.logger.Info("[Migrating...User]")

	if err := r.DB.AutoMigrate(&models.User{}); err != nil {
		r.logger.Error("[Migration failed...User]")
		return err
	}
	r.logger.Info("[Migrating...order]")
	if err := r.DB.AutoMigrate(&models.Order{}); err != nil {
		r.logger.Error("[Migration failed...order]")
		return err
	}
	return nil
}

// ExistsByEmail checks if the user exists by email
func (r *Repository) ExistsByEmail(email string) (bool, error) {
	r.logger.Info("[UserRepository...Exists]")

	users := make([]models.User, 0, 1)
	query := r.DB.Where("email = ?", email).Limit(1).Find(&users)

	return query.RowsAffected > 0, query.Error
}

// get all order
func (r *Repository) GetAllOrder() (order []OrderSerializer, err error) {
	query := r.Model(&models.Order{}).Select("orders.*, users.first_name").
		Joins("LEFT JOIN users ON users.id = orders.user_id")

	return order, query.Find(&order).Error

}

// get total qunatity for each product

func (r *Repository) GetTotalOrderForProduct(filter OrderGroupListFilter) (order []OrderForProductSerializer, err error) {
	query := r.Model(&models.Order{}).Select("product_id, sum(quantity) as total_quantity").Group("product_id")
	if filter.QuantityAbove != 0 {
		query.Having("sum(quantity) > ?", filter.QuantityAbove) //filter for group
	}
	return order, query.Find(&order).Error

}

// delete order
func (r *Repository) DeleteOrderByID(orderId string) (err error) {
	r.logger.Info("Deleting order by id")

	query := r.Model(&models.Order{})

	return query.Unscoped().Where("id = ?", orderId).Delete(&models.Order{}).Error
}
