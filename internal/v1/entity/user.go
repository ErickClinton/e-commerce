package entity

type User struct {
	ID       uint      `gorm:"primaryKey"`
	Username string    `gorm:"unique;not null"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Role     string    `gorm:"not null"`
	Products []Product `gorm:"foreignKey:UserId"`
	Wallet   Wallet    `gorm:"foreignKey:UserId"`
}
