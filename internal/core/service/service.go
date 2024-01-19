package service

import (
	"github.com/sifer169966/device-interactions/internal/core/domain"
	"github.com/sifer169966/device-interactions/internal/core/port"
)

/*
	|--------------------------------------------------------------------------
	| Application's Business Logic
	|--------------------------------------------------------------------------
	|
	| Here you can implement a business logic  for your application
	|
*/

type Service struct {
	repository port.Repository
}

func New(repository port.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (svc *Service) SomeBusinessLogic(request domain.BusinessLogicRequest) error {
	return svc.repository.SomeFunction()
}
