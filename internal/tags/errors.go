package tags

import (
	"net/http"

	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/pkg/errors"
)

var (
	ErrTagQuery            = errors.New("TAG01")
	ErrTagNotFound         = errors.New("TAG02")
	ErrTagPermissionDenied = errors.New("TAG03")
)

var errorToTextMap = map[error]string{
	ErrTagQuery:            "Tag not found",
	ErrTagNotFound:         "Failed to perform operation on Tag",
	ErrTagPermissionDenied: "Forbidden to perform operation on Tag",
}

var errorToHttpCodeMap = map[error]int{
	ErrTagQuery:            http.StatusNotFound,
	ErrTagNotFound:         http.StatusInternalServerError,
	ErrTagPermissionDenied: http.StatusForbidden,
}

func init() {
	appErrors.AddTexts(errorToTextMap)
	appErrors.AddHttpCodes(errorToHttpCodeMap)
}
