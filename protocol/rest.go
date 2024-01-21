package protocol

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/sifer169966/device-interactions/config"
	"github.com/sifer169966/device-interactions/internal/handler"
	"github.com/sifer169966/device-interactions/pkg/middleware"
	"github.com/sifer169966/go-logger"
)

/*
	|--------------------------------------------------------------------------
	| Application Protocol
	|--------------------------------------------------------------------------
	|
	| Here you can choose which protocol your application wants to interact
	| with the client for instance HTTP, gRPC etc.
	|
*/

// ServeREST ...
func ServeREST() error {
	initApp()
	srv := echo.New()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	mdw := middleware.NewMiddleware(c)
	srv.Use(mdw.HandlePanic())
	srv.Use(mdw.IngressRelay())
	srv.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "device-interactions service is running.")
	})
	v1 := srv.Group("/v1")
	hdl := handler.NewRESTHandler(app.svc.deviceSubmissionService, app.pkg.validator)
	devicesV1 := v1.Group("/devices/interactions")

	// /v1/devices
	{
		devicesV1.POST("", hdl.DeviceSubmission)
	}
	errCh := make(chan error, 1)
	go func() {
		signal := <-c
		logger.Info(fmt.Sprintf("got signal %v, start graceful shut down ...", signal))
		// self protection
		mdw.SetServerUnavailable()
		graceful(srv, errCh)
	}()

	err := srv.Start(":" + config.Get().App.HTTPPort)
	if err != nil && err != http.ErrServerClosed {
		logger.Error("unexpected rest-http server error", "error", err)
		errCh <- err
	}
	err = <-errCh
	if err != nil {
		os.Exit(1)
	}
	return nil
}
