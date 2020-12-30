package jwt

import (
	"github.com/JokeTrue/my-arts/internal/users"
	usersDomain "github.com/JokeTrue/my-arts/internal/users/domain"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

const IdentityKey = "USER_PK"

type loginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func payloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	if user, ok := data.(*usersDomain.User); ok {
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

	return &usersDomain.User{ID: int(rawId), Permissions: permissions}
}

func authenticator(store users.Service) func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var request loginRequest
		if err := c.ShouldBind(&request); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		email := request.Username
		password := request.Password

		user, err := store.GetUserByEmail(email)
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		return user, nil
	}
}

func GetJWTMiddleware(db *sqlx.DB) (*jwt.GinJWTMiddleware, error) {
	secretKey := os.Getenv("SECRET_KEY")
	service := usersDomain.NewService(usersDomain.NewStore(db))
	middleware := &jwt.GinJWTMiddleware{
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator(service),
		MaxRefresh:      time.Hour,
		IdentityKey:     IdentityKey,
		Key:             []byte(secretKey),
	}
	return jwt.New(middleware)
}
