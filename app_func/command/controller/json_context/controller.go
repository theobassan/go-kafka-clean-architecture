package json_context

import (
	"net/http"

	"go-kafka-clean-architecture/app_func/command/router"

	"github.com/go-errors/errors"
)

func ReturnError(json router.JsonContextJSON, err error) error {
	jsonErr := json(http.StatusInternalServerError, err)
	if !errors.Is(err, nil) {
		return errors.Wrap(jsonErr, 1)
	}
	return errors.Wrap(err, 1)
}

func ReturnSuccess(json router.JsonContextJSON, statusCode int, response interface{}) error {

	jsonErr := json(statusCode, response)
	if !errors.Is(jsonErr, nil) {
		return errors.Wrap(jsonErr, 1)
	}
	return nil
}
