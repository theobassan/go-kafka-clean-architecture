package sql_handler

import (
	"database/sql"
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/interfaces/database"
	"go-kafka-clean-architecture/app_func/interfaces/repository/sql_handler/model"
	"go-kafka-clean-architecture/app_func/usecases/repository"

	"github.com/go-errors/errors"
)

func CreateProductTranslatedMySql() func(exec database.SqlHandlerExec) repository.ProductTranslatedRepositoryCreate {
	return func(exec database.SqlHandlerExec) repository.ProductTranslatedRepositoryCreate {
		return func(product entities.Product) (int64, error) {

			result, err := runCreateProductTranslatedMySql()(exec)(product)
			if !errors.Is(err, nil) {
				return 0, errors.Wrap(err, 1)
			}

			id, err := result.LastInsertId()
			if !errors.Is(err, nil) {
				return 0, errors.Wrap(err, 1)
			}

			return id, nil
		}
	}
}

func runCreateProductTranslatedMySql() func(exec database.SqlHandlerExec) func(product entities.Product) (sql.Result, error) {
	return func(exec database.SqlHandlerExec) func(product entities.Product) (sql.Result, error) {
		return func(product entities.Product) (sql.Result, error) {
			const query = `
				INSERT INTO
					products_translated(external_id, type, name)
				VALUES
					(?, ?, ?)
				`

			result, err := exec(query, product.ID, product.Type, product.Name)
			if !errors.Is(err, nil) {
				return nil, errors.Wrap(err, 1)
			}

			return result, nil
		}
	}
}

func FindAllProductsTranslatedMySql() func(queryFunc database.SqlHandlerQuery) repository.ProductTranslatedRepositoryFindAll {
	return func(queryFunc database.SqlHandlerQuery) repository.ProductTranslatedRepositoryFindAll {
		return func() ([]entities.Product, error) {

			rows, err := runFindAllProductsTranslatedMySql()(queryFunc)
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

func runFindAllProductsTranslatedMySql() func(queryFunc database.SqlHandlerQuery) (*sql.Rows, error) {
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
