package middleware

import (
	"os"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/sifer169966/device-interactions/pkg/apperror"
	"github.com/sifer169966/device-interactions/pkg/appresponse"
	"github.com/sifer169966/go-logger"
)

const (
	availableMode  = 0
	unavilableMode = 1
)

type Middleware struct {
	mode         int8
	systemSignal chan os.Signal
}

// Implement your middleware
func NewMiddleware(c chan os.Signal) *Middleware {
	return &Middleware{
		mode:         availableMode,
		systemSignal: c,
	}
}

func (mdw *Middleware) IngressRelay() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if mdw.mode == unavilableMode {
				return appresponse.Error(c, apperror.NewUnavailableError())
			}
			return next(c)
		}
	}
}

func (mdw *Middleware) HandlePanic() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				r := recover()
				if r != nil {
					logger.Error("got panic in goroutine", "error", r)
					mdw.systemSignal <- syscall.SIGTERM
					_ = appresponse.Error(c, apperror.NewInternalServerError())
				}

			}()
			return next(c)
		}
	}
}

func (mdw *Middleware) SetServerUnavailable() {
	mdw.mode = unavilableMode
}
