package users

import (
	"net/http"

	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/pkg/errors"
)

var (
	ErrUserQuery    = errors.New("USR01")
	ErrUserNotFound = errors.New("USR02")
)

var errorToTextMap = map[error]string{
	ErrUserNotFound: "User not found",
	ErrUserQuery:    "Failed to perform operation on User",
}

var errorToHttpCodeMap = map[error]int{
	ErrUserNotFound: http.StatusNotFound,
	ErrUserQuery:    http.StatusInternalServerError,
}

func init() {
	appErrors.AddTexts(errorToTextMap)
	appErrors.AddHttpCodes(errorToHttpCodeMap)
}
