package products

import (
	"net/http"

	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/pkg/errors"
)

var (
	ErrProductQuery            = errors.New("PDT01")
	ErrProductNotFound         = errors.New("PDT02")
	ErrProductUnknownState     = errors.New("PDT03")
	ErrProductPermissionDenied = errors.New("PDT04")
)

var errorToTextMap = map[error]string{
	ErrProductQuery:            "Product not found",
	ErrProductNotFound:         "Failed to perform operation on Product",
	ErrProductUnknownState:     "Unknown Product State",
	ErrProductPermissionDenied: "Forbidden to perform operation on Product",
}

var errorToHttpCodeMap = map[error]int{
	ErrProductQuery:            http.StatusNotFound,
	ErrProductNotFound:         http.StatusInternalServerError,
	ErrProductUnknownState:     http.StatusBadRequest,
	ErrProductPermissionDenied: http.StatusForbidden,
}

func init() {
	appErrors.AddTexts(errorToTextMap)
	appErrors.AddHttpCodes(errorToHttpCodeMap)
}
