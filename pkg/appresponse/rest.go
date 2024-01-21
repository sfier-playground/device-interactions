package appresponse

import (
	"fmt"
	"net/http"

	"github.com/sifer169966/device-interactions/config"
	"github.com/sifer169966/device-interactions/pkg/apperror"

	"github.com/labstack/echo/v4"
)

type responseSuccess struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var success = "Success"

func Success(c echo.Context, resp interface{}) error {
	return c.JSON(http.StatusOK, returnSuccess(resp))
}

func NoContent(c echo.Context, resp interface{}) error {
	return c.JSON(http.StatusNoContent, returnSuccess(resp))
}

func returnSuccess(resp interface{}) responseSuccess {
	return responseSuccess{
		Code:    fmt.Sprintf("%s-20001", config.Get().App.CodeName),
		Message: success,
		Data:    resp,
	}
}

type responseError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func Error(c echo.Context, err error) error {
	e, ok := apperror.IsAppError(err)
	if ok {
		return c.JSON(HTTPStatus[e.Code], returnError(int(e.Code), string(e.Message), e.Errors))
	}
	return c.JSON(http.StatusInternalServerError, returnError(int(apperror.ErrInternalServerError), err.Error(), nil))
}

func returnError(code int, message string, errs interface{}) responseError {
	if errs == nil {
		errs = []apperror.ErrorModel{}
	}
	return responseError{
		Code:    fmt.Sprintf("%s-%d", config.Get().App.CodeName, code),
		Message: message,
		Errors:  errs,
	}
}
