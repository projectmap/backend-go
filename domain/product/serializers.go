package product

type ProductSerializer struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	ProductName string `json:"product_name"`
	Price       int16  `json:"price"`
	Quantity    int16  `json:"quantity"`
	ProductType string `json:"product_type"`
}

type ProductListFilter struct {
	Search      string
	ProductType string
	MinPrice    int
	MaxPrice    int
}
