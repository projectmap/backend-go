package product

type ProductSerializer struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	ProductName string `json:"product_name"`
	Price       int16  `json:"price"`
	Quantity    int16  `json:"quantity"`
}
