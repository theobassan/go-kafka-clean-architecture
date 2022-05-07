package sql_handler

import (
	"database/sql"
	"go-kafka-clean-architecture/app_func/entities"
	"go-kafka-clean-architecture/app_func/interfaces/database"
	"go-kafka-clean-architecture/app_func/interfaces/repository/sql_handler/model"
	"go-kafka-clean-architecture/app_func/usecases/repository"

	"github.com/go-errors/errors"
)

func CreateProductMySql() func(exec database.SqlHandlerExec) repository.ProductRepositoryCreate {
	return func(exec database.SqlHandlerExec) repository.ProductRepositoryCreate {
		return func(product entities.Product) (int64, error) {

			result, err := runCreateProductMySql()(exec)(product)
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

func runCreateProductMySql() func(exec database.SqlHandlerExec) func(product entities.Product) (sql.Result, error) {
	return func(exec database.SqlHandlerExec) func(product entities.Product) (sql.Result, error) {
		return func(product entities.Product) (sql.Result, error) {
			const query = `
					INSERT INTO
						products(external_id, type, name)
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

func FindAllProductsMySql() func(queryFunc database.SqlHandlerQuery) repository.ProductRepositoryFindAll {
	return func(queryFunc database.SqlHandlerQuery) repository.ProductRepositoryFindAll {
		return func() ([]entities.Product, error) {

			rows, err := runFindAllProductsMySql()(queryFunc)
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

func runFindAllProductsMySql() func(queryFunc database.SqlHandlerQuery) (*sql.Rows, error) {
	return func(queryFunc database.SqlHandlerQuery) (*sql.Rows, error) {
		const query = `
			SELECT
				external_id,
				type,
				name
			FROM
				products
		`
		rows, err := queryFunc(query)
		if !errors.Is(err, nil) {
			return nil, errors.Wrap(err, 1)
		}

		return rows, nil
	}
}
