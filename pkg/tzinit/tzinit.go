package tzinit

import (
	"github.com/JokeTrue/my-arts/pkg/utils"
	"os"
	"time"
)

func init() {
	utils.Try(os.Setenv("TZ", time.UTC.String()))
}