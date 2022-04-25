package sql_handler

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/usecases/repository"
)

type productRepository struct {
	sqlHandler database.SQLHandler
}

func NewProductRepository(sqlHandler database.SQLHandler) repository.ProductRepository {
	return &productRepository{sqlHandler}
}

func (repository *productRepository) Create(product *entities.Product) (*int64, error) {
	const query = `
		INSERT INTO
			products(external_id, type, name)
		VALUES
			(?, ?, ?)
	`

	result, err := repository.sqlHandler.Exec(query, product.ID, product.Type, product.Name)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (repository *productRepository) FindAll() ([]*entities.Product, error) {
	const query = `
		SELECT
			external_id,
			type,
			name
		FROM
			products
	`
	rows, err := repository.sqlHandler.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*entities.Product{}
	for rows.Next() {
		var externalId int64
		var productType string
		var name string
		if err = rows.Scan(&externalId, &productType, &name); err != nil {
			return nil, err
		}
		product := &entities.Product{
			ID:   &externalId,
			Type: &productType,
			Name: &name,
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
