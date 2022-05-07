package sql_gorm

import (
	"github.com/go-errors/errors"

	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/interfaces/repository/sql_gorm/model"
	"go-kafka-clean-architecture/app/usecases/repository"
)

type productTranslatedRepository struct {
	sqlGorm database.SqlGorm
}

func NewProductTranslatedRepository(sqlGorm database.SqlGorm) repository.ProductTranslatedRepository {
	return &productTranslatedRepository{sqlGorm}
}

func (repository *productTranslatedRepository) FindAll() ([]*entities.Product, error) {
	modelProducts := []*model.ProductTranslated{}
	err := repository.sqlGorm.Find(&modelProducts).Error
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	products := []*entities.Product{}
	for _, modelProduct := range modelProducts {
		product := &entities.Product{
			ID:   modelProduct.ExternalID,
			Type: modelProduct.Type,
			Name: modelProduct.Name,
		}
		products = append(products, product)
	}

	return products, nil
}

func (repository *productTranslatedRepository) Create(product *entities.Product) (*int64, error) {
	modelProduct := model.ProductTranslated{
		ExternalID: product.ID,
		Type:       product.Type,
		Name:       product.Name,
	}

	err := repository.sqlGorm.Create(&modelProduct).Error
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return modelProduct.ID, nil
}
