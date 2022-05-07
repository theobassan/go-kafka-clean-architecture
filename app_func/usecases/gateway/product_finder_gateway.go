package gateway

import (
	"go-kafka-clean-architecture/app_func/entities"
)

type ProductFinderGatewayFindById func(id int64) (entities.Product, error)
