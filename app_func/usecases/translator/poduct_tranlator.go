package translator

import (
	"go-kafka-clean-architecture/app_func/entities"
)

type ProductTranslatorTranslate func(product entities.Product) entities.Product
