package http_context

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-errors/errors"
)

func ReturnError(responseWriter http.ResponseWriter, err error) error {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusInternalServerError)

	encodeErr := json.NewEncoder(responseWriter).Encode(err)
	if !errors.Is(encodeErr, nil) {
		return errors.Wrap(encodeErr, 1)
	}
	return errors.Wrap(err, 1)
}

func ReturnSuccess(responseWriter http.ResponseWriter, statusCode int, resposne any) error {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)

	encodeErr := json.NewEncoder(responseWriter).Encode(resposne)
	if !errors.Is(encodeErr, nil) {
		return errors.Wrap(encodeErr, 1)
	}
	return nil
}

func GetIntParamFromRequest(request *http.Request, paramKey string) (int64, error) {
	param := request.URL.Query().Get(paramKey)
	paramInt, err := strconv.ParseInt(param, 10, 64)
	if !errors.Is(err, nil) {
		return 0, errors.Wrap(err, 1)
	}
	return paramInt, nil
}
