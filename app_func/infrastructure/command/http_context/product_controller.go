package http_context

import "go-kafka-clean-architecture/app_func/command/router"

type ProductControllerCreate func(responseWriterFunc router.HttpContextResponseWriter, requestFunc router.HttpContextRequest) error
type ProductControllerFindAll func(responseWriterFunc router.HttpContextResponseWriter) error
type ProductControllerGet func(responseWriterFunc router.HttpContextResponseWriter, requestFunc router.HttpContextRequest) error
