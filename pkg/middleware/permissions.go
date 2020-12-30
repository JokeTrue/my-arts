package middleware

import (
	appErrors "github.com/JokeTrue/my-arts/pkg/errors"
	perms "github.com/JokeTrue/my-arts/pkg/permissions"
	gJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"strings"
)

func CheckPermissionsMiddleware(handler func(c *gin.Context), permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if permissions == nil {
			handler(c)
			return

		}
		var userPermissions []string

		claims := gJwt.ExtractClaims(c)
		if claimsPerms, ok := claims["perms"]; ok && claimsPerms != nil {
			userPermissions = strings.Split(claimsPerms.(string), ",")
		}

		if !perms.HasPermission(userPermissions, permissions) {
			appErrors.JSONError(c, appErrors.ErrMissingPermissions, permissions)
			return
		}

		handler(c)
	}
}
