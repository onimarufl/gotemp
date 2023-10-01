package service

import (
	"gotemp/model"
	"gotemp/repository"
)

type Service interface {
	GetMessage(message string) string
	PostService(req model.Request) (model.Request, error)
	InsertData(req repository.InsertDataRequest) error
	InquiryDataByID(id int) (repository.User, error)
	InquiryAllData() ([]repository.User, error)
	UpdateDataByID(req repository.UpdateDataRequest) (repository.UpdateDataRequest, error)
	DeleteDataByID(id int) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) service {
	return service{repository: repository}
}

func (s service) GetMessage(message string) string {

	return "Get " + message
}

func (s service) PostService(req model.Request) (model.Request, error) {

	req.Firstname = "Get " + req.Firstname
	req.Lastname = "Get " + req.Lastname

	return req, nil
}

func (s service) InsertData(req repository.InsertDataRequest) error {

	err := s.repository.InsertData(req)
	if err != nil {
		return err
	}

	return nil
}

func (s service) InquiryDataByID(id int) (repository.User, error) {

	res, err := s.repository.InquiryDataByID(id)
	if err != nil {
		return repository.User{}, err
	}

	return res, nil
}

func (s service) InquiryAllData() ([]repository.User, error) {

	res, err := s.repository.InquiryAllData()
	if err != nil {
		return []repository.User{}, err
	}

	return res, nil
}

func (s service) UpdateDataByID(req repository.UpdateDataRequest) (repository.UpdateDataRequest, error) {
	dataUpdateResp, err := s.repository.UpdateDataByID(req)
	if err != nil {
		return repository.UpdateDataRequest{}, err
	}

	return dataUpdateResp, nil
}

func (s service) DeleteDataByID(id int) error {
	err := s.repository.DeleteDataByID(id)
	if err != nil {
		return err
	}

	return nil
}
