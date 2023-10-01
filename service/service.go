package service

import (
	"gotemp/model"
)

type Service interface {
	GetMessage(message string) string
	PostService(req model.Request) (model.Request, error)
}

type service struct {
}

func NewService() service {
	return service{}
}

func (s service) GetMessage(message string) string {

	return "Get " + message
}

func (s service) PostService(req model.Request) (model.Request, error) {

	req.Firstname = "Get " + req.Firstname
	req.Lastname = "Get " + req.Lastname

	return req, nil
}
