package friendship

import (
	"net/http"

	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/pkg/errors"
)

var (
	ErrFriendshipQuery                = errors.New("FSR01")
	ErrFriendshipNotFound             = errors.New("FSR02")
	ErrFriendshipPermissionDenied     = errors.New("FSR03")
	ErrFriendshipUnknownAction        = errors.New("FSR04")
	ErrFriendshipRequestAlreadyExists = errors.New("FSR05")
	ErrFriendshipAlreadyExists        = errors.New("FSR06")
)

var errorToTextMap = map[error]string{
	ErrFriendshipNotFound:             "Friendship not found",
	ErrFriendshipQuery:                "Failed to perform operation on FriendshipRequest",
	ErrFriendshipPermissionDenied:     "Forbidden to perform operation on FriendshipRequest",
	ErrFriendshipUnknownAction:        "Unknown action for FriendshipRequest",
	ErrFriendshipRequestAlreadyExists: "Friendship Request already exists",
	ErrFriendshipAlreadyExists:        "Friendship already exists",
}

var errorToHttpCodeMap = map[error]int{
	ErrFriendshipNotFound:             http.StatusNotFound,
	ErrFriendshipQuery:                http.StatusInternalServerError,
	ErrFriendshipPermissionDenied:     http.StatusForbidden,
	ErrFriendshipUnknownAction:        http.StatusBadRequest,
	ErrFriendshipRequestAlreadyExists: http.StatusOK,
	ErrFriendshipAlreadyExists:        http.StatusBadRequest,
}

func init() {
	appErrors.AddTexts(errorToTextMap)
	appErrors.AddHttpCodes(errorToHttpCodeMap)
}
