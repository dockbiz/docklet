package src

import (
	"golang.org/x/net/context"
	"net/http"
)

func IsDevServer() bool {
	return internal.IsDevAppServer()
}

func NewContext(req *http.Request) context.Context {
	return WithContext(context.Background(), req)
}

func WithContext(parent context.Context, req *http.Request) context.Context {
	return internal.WithContext(parent, req)
}

func init() {
}
