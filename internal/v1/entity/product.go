package entity

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Price       float64       `gorm:"not null"`
	UserId      uint          `gorm:"not null"`
	Carts       []CartProduct `gorm:"foreignKey:ProductId"`
}
