package entity

type Wallet struct {
	ID      uint    `gorm:"primaryKey"`
	UserId  uint    `gorm:"unique;not null"`
	Balance float64 `gorm:"default:0"`
}
