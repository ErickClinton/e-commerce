package entity

type Cart struct {
	Id       uint          `gorm:"primaryKey"`
	UserId   uint          `gorm:"not null;unique"`
	Products []CartProduct `gorm:"foreignKey:CartId"`
}
