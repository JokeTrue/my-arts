package users

import (
	"net/http"

	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/pkg/errors"
)

var (
	ErrUserQuery              = errors.New("USR01")
	ErrUserNotFound           = errors.New("USR02")
	ErrUserAlreadyExists      = errors.New("USR03")
	ErrUserPasswordsDontMatch = errors.New("USR04")
)

var errorToTextMap = map[error]string{
	ErrUserNotFound:           "User not found",
	ErrUserQuery:              "Failed to perform operation on User",
	ErrUserAlreadyExists:      "User already exists",
	ErrUserPasswordsDontMatch: "Password dont match",
}

var errorToHttpCodeMap = map[error]int{
	ErrUserNotFound:           http.StatusNotFound,
	ErrUserQuery:              http.StatusInternalServerError,
	ErrUserAlreadyExists:      http.StatusBadRequest,
	ErrUserPasswordsDontMatch: http.StatusBadRequest,
}

func init() {
	appErrors.AddTexts(errorToTextMap)
	appErrors.AddHttpCodes(errorToHttpCodeMap)
}
