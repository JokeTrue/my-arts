package appErrors

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type ErrorResponse struct {
	ErrorCode string      `json:"error_code"`
	ErrorText string      `json:"error"`
	ExtraInfo interface{} `json:"extra"`
}

var (
	ErrMissingParameter   = errors.New("CMN01")
	ErrBadParameter       = errors.New("CMN02")
	ErrMissingPermissions = errors.New("CMN03")
)

var ErrorToTextMap = map[error]string{
	ErrMissingParameter:   "Missing Parameter",
	ErrBadParameter:       "Bad Parameter",
	ErrMissingPermissions: "Missing Permissions",
}

var ErrorToHttpCodeMap = map[error]int{
	ErrMissingParameter:   http.StatusBadRequest,
	ErrBadParameter:       http.StatusBadRequest,
	ErrMissingPermissions: http.StatusForbidden,
}

func AddTexts(texts map[error]string) {
	for k, v := range texts {
		ErrorToTextMap[k] = v
	}
}

func AddHttpCodes(codes map[error]int) {
	for k, v := range codes {
		ErrorToHttpCodeMap[k] = v
	}
}

func JSONError(c *gin.Context, err error, extra interface{}) {
	errorResponse := ErrorResponse{
		ExtraInfo: extra,
		ErrorText: "unknown",
		ErrorCode: err.Error(),
	}
	if errorText := ErrorToTextMap[err]; errorText != "" {
		errorResponse.ErrorText = errorText
	}

	httpCode := http.StatusInternalServerError
	if code := ErrorToHttpCodeMap[err]; code != 0 {
		httpCode = code
	}

	c.JSON(httpCode, errorResponse)
}
