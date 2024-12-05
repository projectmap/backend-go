package user

type OrderSerializer struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
}
