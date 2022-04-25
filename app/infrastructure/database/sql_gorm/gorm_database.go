package sql_gorm

import (
	"fmt"
	"go-kafka-clean-architecture/app/interfaces/database"

	"github.com/jinzhu/gorm"
)

func NewGormDatabase(dialect string, source string) (database.SQLGorm, error) {

	gormDB, err := gorm.Open(dialect, source)
	if err != nil {
		fmt.Println(1)
		return nil, err
	}
	err = gormDB.DB().Ping()
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
