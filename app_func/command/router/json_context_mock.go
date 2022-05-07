package router

import (
	"encoding/json"
	"net/http"
)

type JsonContextMock struct {
	responseWriter http.ResponseWriter
	request        *http.Request
}

func NewJsonContextMock(responseWriter http.ResponseWriter, request *http.Request) JsonContextMock {
	return JsonContextMock{responseWriter, request}
}

func (m JsonContextMock) JSON(code int, i interface{}) error {
	m.responseWriter.Header().Set("Content-Type", "application/json")
	m.responseWriter.WriteHeader(code)

	return json.NewEncoder(m.responseWriter).Encode(i)
}

func (m JsonContextMock) Bind(v any) error {
	return json.NewDecoder(m.request.Body).Decode(v)
}

func (m JsonContextMock) Query(key string) string {
	return m.request.URL.Query().Get(key)
}
