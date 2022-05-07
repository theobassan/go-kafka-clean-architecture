package rest_api

import (
	"go-kafka-clean-architecture/app/interfaces/api"
	"net/http"
)

type httpApi struct {
	baseURL string
}

func NewHttpApi(baseURL string) api.RestApi {
	return &httpApi{baseURL}
}

func (httpApi *httpApi) Get(url string) (*http.Response, error) {
	return http.Get(httpApi.baseURL + "/" + url)
}
