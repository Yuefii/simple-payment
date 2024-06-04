package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
}

type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
	User      User    `gorm:"foreignKey:UserID"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&Product{},
		&Order{},
	)
}
