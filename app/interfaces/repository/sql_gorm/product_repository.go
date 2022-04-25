package sql_gorm

import (
	"errors"

	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/interfaces/repository/sql_gorm/model"
	"go-kafka-clean-architecture/app/usecases/repository"
)

type productRepository struct {
	sqlGorm database.SQLGorm
}

func NewProductRepository(sqlGorm database.SQLGorm) repository.ProductRepository {
	return &productRepository{sqlGorm}
}

func (repository *productRepository) FindAll() ([]*entities.Product, error) {
	modelProducts := []*model.Product{}
	if err := repository.sqlGorm.Find(&modelProducts).Error; !errors.Is(err, nil) {
		return nil, err
	}

	products := []*entities.Product{}
	for _, modelProduct := range modelProducts {
		product := &entities.Product{
			ID:   modelProduct.ID,
			Type: modelProduct.Type,
			Name: modelProduct.Name,
		}
		products = append(products, product)
	}

	return products, nil
}

func (repository *productRepository) Create(product *entities.Product) (*int64, error) {
	modelProduct := model.Product{
		ExternalID: product.ID,
		Type:       product.Type,
		Name:       product.Name,
	}

	if err := repository.sqlGorm.Create(&modelProduct).Error; !errors.Is(err, nil) {
		return nil, err
	}

	return modelProduct.ID, nil
}
