package api

import "net/http"

type RestApi interface {
	Get(url string) (*http.Response, error)
}
