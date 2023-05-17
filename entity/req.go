package entity

import "net/http"

type Request struct {
	Body    []byte
	Headers http.Header
	Method  string
	URL     string
}
