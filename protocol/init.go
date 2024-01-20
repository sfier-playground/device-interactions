package protocol

import (
	"log"

	"github.com/sifer169966/device-interactions/config"
	"github.com/sifer169966/device-interactions/infrastructure"
	"github.com/sifer169966/device-interactions/internal/core/service"
	"github.com/sifer169966/device-interactions/internal/repository"
	"github.com/sifer169966/device-interactions/pkg/flags"
	"github.com/sifer169966/device-interactions/pkg/validators"
	"github.com/sifer169966/go-logger"
)

var app *application

type application struct {
	// svc stand for service
	svc services
	// pkg stand for package
	pkg packages
}

type services struct {
	deviceSubmissionService *service.DeviceSubmissionService
}

type packages struct {
	validator *validators.Validator
}

func initApp() {
	err := logger.Init(logger.Config{
		ServiceName:    "device-interactions",
		ServiceVersion: flags.Version,
		Level:          "INFO",
		Format:         "json",
	})
	if err != nil {
		log.Fatalf("cound not not init logger: %v", err)
	}
	err = config.Init()
	if err != nil {
		logger.Fatal("could not parsing environment", "error", err)
	}
	// prepare packages
	vld, err := validators.New()
	if err != nil {
		log.Fatalf("could not create the validator instance: %v", err)
	}
	packages := packages{
		validator: vld,
	}
	dbConn := infrastructure.NewPostgreSQLConnection(infrastructure.PostgreSQLConfig{})
	deviceInteractionsRepo := repository.NewDeviceInteractions(dbConn)
	app = &application{
		svc: services{
			deviceSubmissionService: service.NewDeviceSubmission(deviceInteractionsRepo),
		},
		pkg: packages,
	}
}
