package sql_handler

import (
	"database/sql"
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/interfaces/database"
	"go-kafka-clean-architecture/app_func/interfaces/repository/sql_handler/model"
	"go-kafka-clean-architecture/app_func/usecases/repository"

	"github.com/go-errors/errors"
)

func CreateProductTranslatedPostgres() func(queryRow database.SqlHandlerQueryRow) repository.ProductTranslatedRepositoryCreate {
	return func(queryRow database.SqlHandlerQueryRow) repository.ProductTranslatedRepositoryCreate {
		return func(product entities.Product) (int64, error) {

			row := runCreateProductTranslatedPostgres()(queryRow)(product)

			var id int64
			err := row.Scan(&id)
			if !errors.Is(err, nil) {
				return int64(0), errors.Wrap(err, 1)
			}

			return id, nil
		}
	}
}

func runCreateProductTranslatedPostgres() func(queryRow database.SqlHandlerQueryRow) func(product entities.Product) *sql.Row {
	return func(queryRow database.SqlHandlerQueryRow) func(product entities.Product) *sql.Row {
		return func(product entities.Product) *sql.Row {
			const query = `
				INSERT INTO
					products_translated(external_id, type, name)
				VALUES
					($1, $2, $3)
				RETURNING
					id
			`

			row := queryRow(query, product.ID, product.Type, product.Name)

			return row
		}
	}
}

func FindAllProductsTranslatedPostgres() func(queryFunc database.SqlHandlerQuery) repository.ProductTranslatedRepositoryFindAll {
	return func(queryFunc database.SqlHandlerQuery) repository.ProductTranslatedRepositoryFindAll {
		return func() ([]entities.Product, error) {

			rows, err := runFindAllProductsTranslatedPostgres()(queryFunc)
			if !errors.Is(err, nil) {
				return nil, errors.Wrap(err, 1)
			}
			defer rows.Close()

			products, err := model.MapProductsFromRows(rows)
			if !errors.Is(err, nil) {
				return nil, errors.Wrap(err, 1)
			}

			err = rows.Err()
			if !errors.Is(err, nil) {
				return nil, errors.Wrap(err, 1)
			}

			return products, nil
		}
	}
}

func runFindAllProductsTranslatedPostgres() func(queryFunc database.SqlHandlerQuery) (*sql.Rows, error) {
	return func(queryFunc database.SqlHandlerQuery) (*sql.Rows, error) {
		const query = `
			SELECT
				external_id,
				type,
				name
			FROM
				products_translated
		`

		rows, err := queryFunc(query)
		if !errors.Is(err, nil) {
			return nil, errors.Wrap(err, 1)
		}

		return rows, nil
	}
}
