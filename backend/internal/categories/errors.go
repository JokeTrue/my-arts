package categories

import (
	"net/http"

	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/pkg/errors"
)

var (
	ErrCategoryQuery            = errors.New("CAT01")
	ErrCategoryNotFound         = errors.New("CAT02")
	ErrCategoryPermissionDenied = errors.New("CAT03")
)

var errorToTextMap = map[error]string{
	ErrCategoryQuery:            "Category not found",
	ErrCategoryNotFound:         "Failed to perform operation on Category",
	ErrCategoryPermissionDenied: "Forbidden to perform operation on Category",
}

var errorToHttpCodeMap = map[error]int{
	ErrCategoryQuery:            http.StatusNotFound,
	ErrCategoryNotFound:         http.StatusInternalServerError,
	ErrCategoryPermissionDenied: http.StatusForbidden,
}

func init() {
	appErrors.AddTexts(errorToTextMap)
	appErrors.AddHttpCodes(errorToHttpCodeMap)
}
