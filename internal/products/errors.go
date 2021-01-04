package products

import (
	"net/http"

	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/pkg/errors"
)

var (
	ErrProductQuery        = errors.New("PDT01")
	ErrProductNotFound     = errors.New("PDT02")
	ErrUnknownProductState = errors.New("PDT03")
)

var errorToTextMap = map[error]string{
	ErrProductQuery:        "Product not found",
	ErrProductNotFound:     "Failed to perform operation on Product",
	ErrUnknownProductState: "Unknown Product State",
}

var errorToHttpCodeMap = map[error]int{
	ErrProductQuery:        http.StatusNotFound,
	ErrProductNotFound:     http.StatusInternalServerError,
	ErrUnknownProductState: http.StatusBadRequest,
}

func init() {
	appErrors.AddTexts(errorToTextMap)
	appErrors.AddHttpCodes(errorToHttpCodeMap)
}
