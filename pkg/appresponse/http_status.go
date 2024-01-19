package appresponse

import (
	"net/http"

	"github.com/sifer169966/device-interactions/pkg/apperror"
)

var (
	HTTPStatus = map[apperror.ErrorCode]int{
		apperror.ErrBadRequestCode:      http.StatusBadRequest,
		apperror.ErrInternalServerError: http.StatusInternalServerError,
		apperror.ErrServerUnavailable:   http.StatusServiceUnavailable,
	}
)
