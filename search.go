package main

import (
	"regexp"
	"strings"
)

var QueryRegexp regexp.Regexp = *regexp.MustCompile(`(/[\w\d]+)(\W+.*)?`)

// The query is divided into 2 parts
// 1) ALIASQUERY - used to search for the link.
// 2) TAIL - appended as-is to the resulting link.
//
// [Example]
// Request: https://this-servers-domain.com/foo/fighters?n=69
// ALIASQUERY: /foo
// TAIL: /fighters?n=69
// Result: https://resulting-link.com/fighters?n=69
func searchNearestLink(query string) (x string, exists bool) {
	if len(query) == 0 {
		return "", false
	}

	queryParts := QueryRegexp.FindStringSubmatch(query)
	if queryParts == nil {
		return "", false
	}

	aliasQuery := queryParts[1]

	var tail string
	if len(queryParts) > 1 {
		tail = queryParts[2]
	}

	// Check exact match
	if link, exists := config.Aliases[aliasQuery]; exists {
		return (link + tail), true
	}

	// Search as prefix
	for alias, link := range config.Aliases {
		if strings.HasPrefix(alias, aliasQuery) {
			return (link + tail), true
		}
	}

	// Search as suffix
	noSlashKey := aliasQuery[1:]
	for alias, link := range config.Aliases {
		if strings.HasSuffix(alias, noSlashKey) {
			return (link + tail), true
		}
	}

	return "", false
}
