package models

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"

	"gorm.io/gorm"
)

// order model
type Order struct {
	gorm.Model

	Quantity int16 `json:"quantity"`
	// Foreign keys

	UserID    uint `json:"user_id"`    // Foreign key for User
	ProductID uint `json:"product_id"` // Foreign key for Product

	// Relationships
	User    User    `json:"user" gorm:"foreignKey:UserID"`       // Belongs to User
	Product Product `json:"product" gorm:"foreignKey:ProductID"` // Belongs to Product
}

func (*Order) TableName() string {
	return "orders"
}
