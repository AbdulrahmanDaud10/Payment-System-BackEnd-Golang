package api

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;" json:"updated_at"`
}

// User --> Model for User table
type User struct {
	gorm.Model
	UserID      string `gorm:"column:user_id;" json:"user_id"`
	Email       string `gorm:"column:email;" json:"email"`
	PhoneNumber string `gorm:"column:phone_number;" json:"phone_number"`
}

// Product --> Model for Product table
type Product struct {
	gorm.Model
	ProductID string `gorm:"column:product_id;" json:"product_id"`
	Code      string `gorm:"column:code;" json:"code"`
	Price     uint   `gorm:"column:price;" json:"price"`
}

// BankAccount --> Model for BankAccount table
type BankAccount struct {
	gorm.Model
	BankAccountID string  `gorm:"column:bank_account_id;" json:"bank_account_id"`
	Balance       float64 `gorm:"column:balance;" json:"balance"`
}
