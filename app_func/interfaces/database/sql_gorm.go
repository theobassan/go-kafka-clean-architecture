package database

import "github.com/jinzhu/gorm"

type SqlGormFind func(out interface{}, where ...interface{}) *gorm.DB
type SqlGormCreate func(value interface{}) *gorm.DB
