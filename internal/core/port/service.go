package port

import (
	"github.com/sifer169966/device-interactions/internal/core/domain"
)

type ServiceDeviceInteractor interface {
	Submit(in domain.DeviceSubmission) error
}
