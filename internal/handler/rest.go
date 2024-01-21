package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sifer169966/device-interactions/internal/core/port"
	"github.com/sifer169966/device-interactions/pkg/apperror"
	"github.com/sifer169966/device-interactions/pkg/appresponse"
	"github.com/sifer169966/go-logger"
)

type Validator interface {
	ValidateStruct(inf interface{}) error
}

type RESTHandler struct {
	svc port.ServiceDeviceInteractor
	vld Validator
}

func NewRESTHandler(svc port.ServiceDeviceInteractor, vld Validator) *RESTHandler {
	return &RESTHandler{
		svc: svc,
		vld: vld,
	}
}

// DeviceSubmission ...
// handling the interaction from devices
func (hdl *RESTHandler) DeviceSubmission(c echo.Context) error {
	var req DeviceSubmissionRequest
	err := c.Bind(&req)
	if err != nil {
		logger.Info("could not decode the request payload", "info", err)
		return appresponse.Error(c, apperror.NewBadRequestError())
	}
	err = hdl.vld.ValidateStruct(req)
	if err != nil {
		logger.Info("incorrect payload format", "info", err)
		return appresponse.Error(c, err)
	}
	err = hdl.svc.Submit(req.toDeviceSubmissionDomain())
	if err != nil {
		return appresponse.Error(c, err)
	}
	return appresponse.NoContent(c, nil)
}
