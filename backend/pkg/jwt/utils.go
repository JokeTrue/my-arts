package jwt

import (
	"github.com/JokeTrue/my-arts/internal/models"
	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	"github.com/gin-gonic/gin"
)

func GetCurrentUserID(c *gin.Context) (int, error) {
	rawUser, ok := c.Get(IdentityKey)
	if !ok {
		return 0, appErrors.ErrAuthenticationFailed
	}

	jwtUser, ok := rawUser.(*models.User)
	if !ok {
		return 0, appErrors.ErrAuthenticationFailed
	}

	return jwtUser.ID, nil
}
