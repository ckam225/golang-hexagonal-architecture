package server

type Server interface {
	Start()
	Test() any
}
