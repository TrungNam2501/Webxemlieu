package models

// Item là model đơn giản đại diện cho một mặt hàng.
type Item struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
