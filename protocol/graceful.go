package protocol

import (
	"context"
	"time"

	"github.com/sifer169966/go-logger"

	"github.com/labstack/echo/v4"
)

func graceful(srv *echo.Echo, errCh chan error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		logger.Error("could not shut down server", "error", err)
		errCh <- err
	}
	errCh <- nil
}
