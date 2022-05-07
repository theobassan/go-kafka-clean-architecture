package api

import "net/http"

type RestApiGet func(url string) (*http.Response, error)
