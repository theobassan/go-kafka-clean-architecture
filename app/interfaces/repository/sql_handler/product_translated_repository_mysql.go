package sql_handler

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/usecases/repository"

	"github.com/go-errors/errors"
)

type productTranslatedRepositoryMySql struct {
	sqlHandler database.SQLHandler
}

func NewProductTranslatedRepositoryMySql(sqlHandler database.SQLHandler) repository.ProductTranslatedRepository {
	return &productTranslatedRepositoryMySql{sqlHandler}
}

func (repository *productTranslatedRepositoryMySql) Create(product *entities.Product) (*int64, error) {
	const query = `
		INSERT INTO
			products_translated(external_id, type, name)
		VALUES
			(?, ?, ?)
	`

	result, err := repository.sqlHandler.Exec(query, product.ID, product.Type, product.Name)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	id, err := result.LastInsertId()
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return &id, nil
}

func (repository *productTranslatedRepositoryMySql) FindAll() ([]*entities.Product, error) {
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
		return nil, errors.Wrap(err, 1)
	}
	defer rows.Close()

	products := []*entities.Product{}
	for rows.Next() {
		var externalId int64
		var productType string
		var name string
		err = rows.Scan(&externalId, &productType, &name)
		if !errors.Is(err, nil) {
			return nil, errors.Wrap(err, 1)
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
		return nil, errors.Wrap(err, 1)
	}

	return products, nil
}
