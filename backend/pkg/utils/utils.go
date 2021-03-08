package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Contains returns true if target string is present in the strings slice.
// Comparison is case-insensitive.
func Contains(slice []string, lookup string) bool {
	for _, val := range slice {
		if strings.EqualFold(val, lookup) {
			return true
		}
	}
	return false
}

func Try(err error) {
	if err != nil {
		panic(err)
	}
}

func GetOffsetLimit(c *gin.Context) (int, int, error) {
	rawOffset := c.Query("offset")
	offset, err := strconv.Atoi(rawOffset)
	if err != nil {
		return 0, 0, errors.Wrap(err, "failed to parse offset")
	}

	rawLimit := c.Query("limit")
	limit, err := strconv.Atoi(rawLimit)
	if err != nil {
		return 0, 0, errors.Wrap(err, "failed to parse limit")
	}

	return offset, limit, nil
}

func GetReadDatabase(databases []*sqlx.DB) *sqlx.DB {
	rand.Seed(time.Now().Unix())
	return databases[rand.Intn(len(databases))]
}
