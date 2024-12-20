package user

type OrderSerializer struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Quantity  int16  `json:"quantity"`
	UserID    uint   `json:"user_id"`
	ProductID uint   `json:"product_id"`
	FirstName string `json:"first_name"`
}

type OrderForProductSerializer struct {
	ProductID     uint `json:"product_id"`
	TotalQuantity uint `json:"total_quantity"`
}

type OrderGroupListFilter struct {
	QuantityAbove int `json:"quantity_above"`
}
