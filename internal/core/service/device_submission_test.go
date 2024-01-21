package service_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/sifer169966/device-interactions/internal/core/domain"
	"github.com/sifer169966/device-interactions/internal/core/service"
	. "github.com/sifer169966/device-interactions/mocks/repository"
	"github.com/sifer169966/go-logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	ctx                          context.Context
	deviceSubmissionService      *service.DeviceSubmissionService
	deviceInteractionsRepository *MockDeviceInteractionsRepository
}

func TestService(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) SetupTest() {
	_ = logger.Init(logger.Config{})
	ctrl := gomock.NewController(ts.T())
	defer ctrl.Finish()
	ts.ctx = context.Background()
	ts.deviceInteractionsRepository = NewMockDeviceInteractionsRepository(ctrl)
	ts.deviceSubmissionService = service.NewDeviceSubmission(
		ts.deviceInteractionsRepository,
	)
}

var (
	errFoo = errors.New("foo")
)

func (ts *TestSuite) TestSubmit() {
	ts.T().Run("Success_ReturnNil", func(t *testing.T) {
		// Set up test data
		deviceSubmission := domain.DeviceSubmission{
			Timestamp: time.Now(),
			Location:  domain.GeoLocation{Latitude: newDecimal("80"), Longitude: newDecimal("80")}, // Adjust with valid values
			Devices: []domain.Device{
				{DeviceID: uuid.NewString(), Name: "Device1"},
				{DeviceID: uuid.NewString(), Name: "Device2"},
			},
		}
		// Set expectations on the mock repository for CreateMany
		ts.deviceInteractionsRepository.EXPECT().CreateMany(gomock.Any()).DoAndReturn(func(input domain.DeviceSubmission) error {
			// Custom validation for the input data, for example, ensuring valid timestamps, location, etc.
			assert.Equal(ts.T(), deviceSubmission.Timestamp, input.Timestamp)
			assert.Equal(ts.T(), deviceSubmission.Location, input.Location)
			// assert.Condition(ts.T(), func() bool {
			for i := 0; i < len(input.Devices); i++ {
				assert.Equal(ts.T(), deviceSubmission.Devices[i].DeviceID, input.Devices[i].DeviceID)
				assert.Equal(ts.T(), deviceSubmission.Devices[i].Name, input.Devices[i].Name)
				interactionID := input.Devices[i].GetInteractionID()
				assert.NotEmpty(ts.T(), interactionID)
				err := uuid.Validate(interactionID)
				if err != nil {
					ts.Fail(fmt.Sprintf("Should be UUID format, but was %v", interactionID))
				}
			}
			// 	return true
			// })
			return nil
		}).Times(1)
		err := ts.deviceSubmissionService.Submit(deviceSubmission)
		ts.NoError(err)
	})

	ts.T().Run("ErrorWhileCreateMany_ReturnError", func(t *testing.T) {
		ts.deviceInteractionsRepository.EXPECT().CreateMany(gomock.Any()).Return(errFoo).Times(1)
		err := ts.deviceSubmissionService.Submit(domain.DeviceSubmission{})
		ts.Error(err)
	})
}

func newDecimal(s string) decimal.Decimal {
	v, err := decimal.NewFromString(s)
	if err != nil {
		panic(err)
	}
	return v
}
