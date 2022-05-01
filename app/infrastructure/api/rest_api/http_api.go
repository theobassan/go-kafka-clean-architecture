package rest_api

import (
	"go-kafka-clean-architecture/app/interfaces/api"
	"net/http"
)

type httpAPI struct {
	baseURL string
}

func NewHttpAPI(baseURL string) api.RestAPI {
	return &httpAPI{baseURL}
}

func (httpAPI *httpAPI) Get(url string) (*http.Response, error) {
	return http.Get(httpAPI.baseURL + "/" + url)
}
