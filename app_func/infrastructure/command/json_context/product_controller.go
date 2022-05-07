package json_context

import "go-kafka-clean-architecture/app_func/command/router"

type ProductControllerCreate func(bind router.JsonContextBind, json router.JsonContextJSON) error
type ProductControllerFindAll func(json router.JsonContextJSON) error
type ProductControllerGet func(query router.JsonContextQuery, json router.JsonContextJSON) error
