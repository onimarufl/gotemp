package service

type Service interface {
	GetMessage(message string) string
}

type service struct {
}

func NewService() service {
	return service{}
}

func (s service) GetMessage(message string) string {

	return "Get " + message
}
