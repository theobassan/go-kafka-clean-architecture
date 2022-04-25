package model

type Product struct {
	ID         *int64  `gorm:"primary_key:id"`
	ExternalID *int64  `gorm:"column:external_id"`
	Type       *string `gorm:"column:type"`
	Name       *string `gorm:"column:name"`
}

func (Product) TableName() string { return "products" }
