package jwt

import (
	"database/sql"
	"os"
	"time"

	"github.com/JokeTrue/my-arts/internal/models"
	"github.com/JokeTrue/my-arts/internal/users"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const IdentityKey = "USER_PK"

type loginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func payloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	if user, ok := data.(*models.User); ok {
		claims[IdentityKey] = user.ID
		claims["perms"] = user.Permissions
	}
	return claims
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	rawId := claims[IdentityKey].(float64)

	var permissions string
	if perms, ok := claims["perms"].(string); ok {
		permissions = perms
	}

	return &models.User{ID: int(rawId), Permissions: sql.NullString{String: permissions, Valid: true}}
}

func authenticator(useCase users.UseCase) func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var request loginRequest
		if err := c.ShouldBind(&request); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		email := request.Username
		password := request.Password

		user, err := useCase.GetUserByEmail(email)
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		return user, nil
	}
}

func GetJWTMiddleware(useCase users.UseCase) (*jwt.GinJWTMiddleware, error) {
	secretKey := os.Getenv("SECRET_KEY")
	week, _ := time.ParseDuration("168h")
	middleware := &jwt.GinJWTMiddleware{
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator(useCase),
		Timeout:         week,
		MaxRefresh:      time.Hour,
		IdentityKey:     IdentityKey,
		Key:             []byte(secretKey),
	}
	return jwt.New(middleware)
}
