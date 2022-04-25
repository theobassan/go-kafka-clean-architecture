package api

import "net/http"

type RestAPI interface {
	Get(url string) (*http.Response, error)
}
