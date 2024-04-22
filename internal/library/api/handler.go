package api

import "github.com/akshay/libraryAssign/internal/library/service"

type Handler struct {
	myService *service.Service
}

func NewHandler(myService *service.Service) *Handler {
	return &Handler{
		myService: myService,
	}
}
