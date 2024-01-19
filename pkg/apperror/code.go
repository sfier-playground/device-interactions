package apperror

type (
	ErrorCode    int
	ErrorMessage string
)

const (
	ErrBadRequestCode      ErrorCode = 40001
	ErrInternalServerError ErrorCode = 50001
	ErrServerUnavailable   ErrorCode = 50003
)

var (
	Message = map[ErrorCode]ErrorMessage{
		ErrBadRequestCode:      "bad request",
		ErrServerUnavailable:   "server is unavailable, please try again later",
		ErrInternalServerError: "internal server error, please try again later",
	}
)
