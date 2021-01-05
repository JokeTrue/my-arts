package reviews

import (
	"net/http"

	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/pkg/errors"
)

var (
	ErrReviewQuery            = errors.New("RVW01")
	ErrReviewNotFound         = errors.New("RVW02")
	ErrReviewPermissionDenied = errors.New("RVW03")
)

var errorToTextMap = map[error]string{
	ErrReviewQuery:            "Review not found",
	ErrReviewNotFound:         "Failed to perform operation on Review",
	ErrReviewPermissionDenied: "Forbidden to perform operation on Review",
}

var errorToHttpCodeMap = map[error]int{
	ErrReviewQuery:            http.StatusNotFound,
	ErrReviewNotFound:         http.StatusInternalServerError,
	ErrReviewPermissionDenied: http.StatusForbidden,
}

func init() {
	appErrors.AddTexts(errorToTextMap)
	appErrors.AddHttpCodes(errorToHttpCodeMap)
}
