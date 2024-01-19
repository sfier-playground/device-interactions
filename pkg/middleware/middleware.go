package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sifer169966/device-interactions/pkg/apperror"
	"github.com/sifer169966/device-interactions/pkg/appresponse"
)

const (
	availableMode  = 0
	unavilableMode = 1
)

type Middleware struct {
	mode int8
}

// Implement your middleware
func NewMiddleware() *Middleware {
	return &Middleware{
		mode: availableMode,
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

func (mdw *Middleware) SetServerUnavailable() {
	mdw.mode = unavilableMode
}
