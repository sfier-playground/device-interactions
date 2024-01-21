package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"github.com/sifer169966/device-interactions/internal/handler"
	. "github.com/sifer169966/device-interactions/mocks/service"
	"github.com/sifer169966/device-interactions/pkg/validators"
	"github.com/sifer169966/go-logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	ctx                     context.Context
	restHandler             *handler.RESTHandler
	serviceDeviceInteractor *MockServiceDeviceInteractor
}

func TestService(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) SetupTest() {
	_ = logger.Init(logger.Config{Level: "FATAL"})
	ctrl := gomock.NewController(ts.T())
	defer ctrl.Finish()
	ts.ctx = context.Background()
	vld, err := validators.New()
	if err != nil {
		ts.T().Error(err)
	}
	ts.serviceDeviceInteractor = NewMockServiceDeviceInteractor(ctrl)
	ts.restHandler = handler.NewRESTHandler(
		ts.serviceDeviceInteractor,
		vld,
	)
}

func (ts *TestSuite) TestDeviceSubmission() {
	ts.T().Run("Success_Response204", func(t *testing.T) {
		e := echo.New()
		requestPayload := handler.DeviceSubmissionRequest{
			Timestamp: time.Now(),
			Location: handler.GeoLocation{
				Latitude:  newDecimal("80"),
				Longitude: newDecimal("90"),
			},
			Devices: []handler.Device{
				{
					ID:   uuid.NewString(),
					Name: "device-foo-bar",
				},
			},
		}
		requestJSON, err := json.Marshal(requestPayload)
		if err != nil {
			ts.T().Error(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/v1/devices/interactions", strings.NewReader(string(requestJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ts.serviceDeviceInteractor.EXPECT().Submit(gomock.Any()).Return(nil).Times(1)
		err = ts.restHandler.DeviceSubmission(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
	ts.T().Run("GotErrorWhileSubmitDevice_Response500", func(t *testing.T) {
		e := echo.New()
		requestPayload := handler.DeviceSubmissionRequest{
			Timestamp: time.Now(),
			Location: handler.GeoLocation{
				Latitude:  newDecimal("80"),
				Longitude: newDecimal("90"),
			},
			Devices: []handler.Device{
				{
					ID:   uuid.NewString(),
					Name: "device-foo-bar",
				},
			},
		}
		requestJSON, err := json.Marshal(requestPayload)
		if err != nil {
			ts.T().Error(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/v1/devices/interactions", strings.NewReader(string(requestJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ts.serviceDeviceInteractor.EXPECT().Submit(gomock.Any()).Return(errors.New("foo")).Times(1)
		err = ts.restHandler.DeviceSubmission(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
	ts.T().Run("RequestWithInvalidPayloadType_Response400WithErrorMessage", func(t *testing.T) {
		e := echo.New()
		requestPayload := map[string]interface{}{
			"timestamp": "bar",
		}
		requestJSON, _ := json.Marshal(requestPayload)
		req := httptest.NewRequest(http.MethodPost, "/v1/devices/interactions", strings.NewReader(string(requestJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := ts.restHandler.DeviceSubmission(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
	testValidationRule(ts)
}

func testValidationRule(ts *TestSuite) {
	defaultRequestPayload := handler.DeviceSubmissionRequest{
		Timestamp: time.Now(),
		Location: handler.GeoLocation{
			Latitude:  newDecimal("-85.05115"),
			Longitude: newDecimal("90"),
		},
		Devices: []handler.Device{
			{
				ID:   uuid.NewString(),
				Name: "device-foo-bar",
			},
		},
	}

	testCases := []struct {
		name        string
		payloadFunc func() interface{}
	}{
		{
			name: "BodyFieldLatitudeLessThanMinimum",
			payloadFunc: func() interface{} {
				requestPayload := defaultRequestPayload
				requestPayload.Location.Latitude = newDecimal("-85.05116")
				return requestPayload
			},
		},
		{
			name: "BodyFieldLatitudeGreaterThanMaximum",
			payloadFunc: func() interface{} {
				requestPayload := defaultRequestPayload
				requestPayload.Location.Latitude = newDecimal("86")
				return requestPayload
			},
		},
		{
			name: "BodyFieldLongitudeLessThanMinimum",
			payloadFunc: func() interface{} {
				requestPayload := defaultRequestPayload
				requestPayload.Location.Longitude = newDecimal("-181")
				return requestPayload
			},
		},
		{
			name: "BodyFieldLongitudeGreaterThanMaximum",
			payloadFunc: func() interface{} {
				requestPayload := defaultRequestPayload
				requestPayload.Location.Longitude = newDecimal("181")
				return requestPayload
			},
		},
		{
			name: "BodyFieldDeviceIDIsNotUUID",
			payloadFunc: func() interface{} {
				requestPayload := defaultRequestPayload
				devices := make([]handler.Device, len(requestPayload.Devices))
				copy(devices, requestPayload.Devices)
				devices[0].ID = "foo"
				requestPayload.Devices = devices
				return requestPayload
			},
		},
		{
			name: "BodyFieldDeviceNameFormatIsInvalid",
			payloadFunc: func() interface{} {
				requestPayload := defaultRequestPayload
				devices := make([]handler.Device, len(requestPayload.Devices))
				copy(devices, requestPayload.Devices)
				devices[0].Name = "foo"
				requestPayload.Devices = devices
				return requestPayload
			},
		},
	}
	defaultExpectedMessage := "Response400"
	for _, test := range testCases {
		payload := test.payloadFunc()
		ts.T().Run(test.name+"_"+defaultExpectedMessage, func(t *testing.T) {
			e := echo.New()
			requestJSON, err := json.Marshal(payload)
			if err != nil {
				ts.T().Error(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/v1/devices/interactions", strings.NewReader(string(requestJSON)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err = ts.restHandler.DeviceSubmission(c)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		})
	}
}

func newDecimal(s string) decimal.Decimal {
	v, err := decimal.NewFromString(s)
	if err != nil {
		panic(err)
	}
	return v
}
