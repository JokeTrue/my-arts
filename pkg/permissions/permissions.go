package permissions

import (
	"github.com/JokeTrue/my-arts/pkg/utils"
)

var AdminPermission = "ADMIN"

func HasPermission(userPermissions, permissions []string) bool {
	for _, perm := range permissions {
		if !utils.Contains(userPermissions, perm) {
			return false
		}
	}
	return true
}
