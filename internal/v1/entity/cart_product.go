package entity

type CartProduct struct {
	CartId    uint    `gorm:"primaryKey"`
	ProductId uint    `gorm:"primaryKey"`
	Quantity  int     `gorm:"not null;default:1"`
	Product   Product `gorm:"foreignKey:ProductId;references:ID"`
}
