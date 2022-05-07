package rest_api

import (
	"go-kafka-clean-architecture/app_func/interfaces/api"
	"net/http"
)

func Get(baseUrl string) api.RestApiGet {
	return func(url string) (*http.Response, error) {
		return http.Get(baseUrl + "/" + url)
	}
}
