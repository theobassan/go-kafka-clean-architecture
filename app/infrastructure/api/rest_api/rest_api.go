package rest_api

import (
	"go-kafka-clean-architecture/app/interfaces/api"
	"net/http"
)

type restAPI struct {
}

func NewRestAPI() api.RestAPI {
	return &restAPI{}
}

func (restAPI *restAPI) Get(url string) (*http.Response, error) {
	return http.Get(url)
}
