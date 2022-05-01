package model

type ProductTranslated struct {
	ID         *int64  `gorm:"primary_key:id"`
	ExternalID *int64  `gorm:"column:external_id"`
	Type       *string `gorm:"column:type"`
	Name       *string `gorm:"column:name"`
}

func (ProductTranslated) TableName() string { return "products_translated" }
