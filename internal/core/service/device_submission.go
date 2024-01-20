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

type DeviceSubmissionService struct {
	dviRepo port.DeviceInteractionsRepository
}

func NewDeviceSubmission(dviRepo port.DeviceInteractionsRepository) *DeviceSubmissionService {
	return &DeviceSubmissionService{
		dviRepo: dviRepo,
	}
}

func (svc *DeviceSubmissionService) Submit(in domain.DeviceSubmission) error {
	in.Timestamp = in.Timestamp.UTC()
	for i := 0; i < len(in.Devices); i++ {
		in.Devices[i].SetInteractionID()
	}
	return svc.dviRepo.CreateMany(in)
}
