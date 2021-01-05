package tzinit

import (
	"os"
	"time"

	"github.com/JokeTrue/my-arts/pkg/utils"
)

func init() {
	utils.Try(os.Setenv("TZ", time.UTC.String()))
}
