package models

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"

	"gorm.io/gorm"
)

// product model
type Product struct {
	gorm.Model

	ProductName string `json:"product_name" gorm:"size:255"`
	Price       int16  `json:"price"`
	Quantity    int16  `json:"quantity"`
}

func (*Product) TableName() string {
	return "products"
}
