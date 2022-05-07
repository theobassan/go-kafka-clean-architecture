package http_context

import "go-kafka-clean-architecture/app_func/command/router"

type ProductTranslatedControllerFindAll func(responseWriterFunc router.HttpContextResponseWriter) error
