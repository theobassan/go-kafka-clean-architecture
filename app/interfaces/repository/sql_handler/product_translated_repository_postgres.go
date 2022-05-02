package sql_handler

import (
	"go-kafka-clean-architecture/app/entities"
	"go-kafka-clean-architecture/app/interfaces/database"
	"go-kafka-clean-architecture/app/usecases/repository"

	"github.com/go-errors/errors"
)

type productTranslatedRepositoryPostgres struct {
	sqlHandler database.SQLHandler
}

func NewProductTranslatedRepositoryPostgres(sqlHandler database.SQLHandler) repository.ProductTranslatedRepository {
	return &productTranslatedRepositoryPostgres{sqlHandler}
}

func (repository *productTranslatedRepositoryPostgres) Create(product *entities.Product) (*int64, error) {
	const query = `
		INSERT INTO
			products_translated(external_id, type, name)
		VALUES
			($1, $2, $3)
		RETURNING
				id
	`

	id := new(int64)
	err := repository.sqlHandler.QueryRow(query, product.ID, product.Type, product.Name).Scan(id)
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	return id, nil
}

func (repository *productTranslatedRepositoryPostgres) FindAll() ([]*entities.Product, error) {
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
