package sql_handler

import (
	"errors"
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/usecases/repository"
)

type productTranslatedRepository struct {
	sqlHandler database.SQLHandler
}

func NewProductTranslatedRepository(sqlHandler database.SQLHandler) repository.ProductTranslatedRepository {
	return &productTranslatedRepository{sqlHandler}
}

func (repository *productTranslatedRepository) Create(product *entities.Product) (*int64, error) {
	const query = `
		INSERT INTO
			products_translated(external_id, type, name)
		VALUES
			(?, ?, ?)
	`

	result, err := repository.sqlHandler.Exec(query, product.ID, product.Type, product.Name)
	if !errors.Is(err, nil) {
		return nil, err
	}

	id, err := result.LastInsertId()
	if !errors.Is(err, nil) {
		return nil, err
	}

	return &id, nil
}

func (repository *productTranslatedRepository) FindAll() ([]*entities.Product, error) {
	const query = `
		SELECT
			external_id,
			type,
			name
		FROM
			products_translated
	`
	rows, err := repository.sqlHandler.Query(query)
	if !errors.Is(err, nil) {
		return nil, err
	}
	defer rows.Close()

	products := []*entities.Product{}
	for rows.Next() {
		var externalId int64
		var productType string
		var name string
		err = rows.Scan(&externalId, &productType, &name)
		if !errors.Is(err, nil) {
			return nil, err
		}
		product := &entities.Product{
			ID:   &externalId,
			Type: &productType,
			Name: &name,
		}
		products = append(products, product)
	}

	err = rows.Err()
	if !errors.Is(err, nil) {
		return nil, err
	}

	return products, nil
}
