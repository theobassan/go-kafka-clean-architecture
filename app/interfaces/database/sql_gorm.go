package database

import "github.com/jinzhu/gorm"

type SQLGorm interface {
	Find(out interface{}, where ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
}
