package service

type MyService interface {
}

// add any services if require in future
type Service struct{}

// initialize service and return its ref
func NewService() *Service {
	return &Service{}
}
