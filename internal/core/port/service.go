package port

import (
	"github.com/sifer169966/device-interactions/internal/core/domain"
)

type ServiceDeviceInteractor interface {
	SomeBusinessLogic(request domain.BusinessLogicRequest) error
}
