package gateway

import (
	"go-kafka-clean-architecture/app_func/entities"
)

type ProductSenderGatewaySend func(product entities.Product) error
