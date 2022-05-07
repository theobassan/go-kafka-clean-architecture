package event_context

import "go-kafka-clean-architecture/app_func/command/router"

type ProductTranslatedControllerCreate func(bind router.EventContextBind, acknowledge router.EventContextAcknowledge) error
