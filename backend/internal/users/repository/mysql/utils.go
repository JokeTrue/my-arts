package mysql

import (
	"strings"
)

func GetSearchFilters(query string) []string {
	var filters []string
	lastNameFilter := "last_name LIKE '%'"
	firstNameFilter := "first_name LIKE '%'"

	tokens := strings.Split(query, " ")
	if query != "" && len(tokens) > 0 {
		lastNameFilter = strings.ReplaceAll(lastNameFilter, "%", tokens[0]+"%")
	}
	if query != "" && len(tokens) > 1 {
		firstNameFilter = strings.ReplaceAll(firstNameFilter, "%", tokens[1]+"%")
	}
	filters = append(filters, lastNameFilter, firstNameFilter)

	return filters
}
