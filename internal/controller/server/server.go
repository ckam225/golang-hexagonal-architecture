package server

import "net/http"

//go:generate mockgen -source=server.go -destination=mocks/mock.go
type Server interface {
	Start() error
	Test(req *http.Request, msTimeout ...int) (*http.Response, error)
}
