package models

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"

	"gorm.io/gorm"
)

// order model
type Order struct {
	gorm.Model

	ProductName string `json:"product_name" gorm:"size:255"`
	Price       int16  `json:"price"`
}

func (*Order) TableName() string {
	return "orders"
}
