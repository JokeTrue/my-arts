package utils

import "strings"

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