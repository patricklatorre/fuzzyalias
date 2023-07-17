package main

import (
	"regexp"
	"strings"
)

var QueryRegexp regexp.Regexp = *regexp.MustCompile(`(/[\w\d]+)(\W+.*)?`)

// The query is divided into 2 parts
// 1) KEY - used to search for the link.
// 2) OPTS - appended as-is to the resulting link.
//
// [Example]
// Request: https://this-servers-domain.com/foo/fighters?n=69
// KEY: /foo
// OPTS: /fighters?n=69
// Result: https://resulting-link.com/fighters?n=69
func searchNearestLink(query string) (x string, exists bool) {
	if len(query) == 0 {
		return "", false
	}

	queryParts := QueryRegexp.FindStringSubmatch(query)
	if queryParts == nil {
		return "", false
	}

	key := queryParts[1]

	var opts string
	if len(queryParts) > 1 {
		opts = queryParts[2]
	}

	// Check exact match
	if link, exists := config.LinkMap[key]; exists {
		return (link + opts), true
	}

	// Search as prefix
	for fullKey, link := range config.LinkMap {
		if strings.HasPrefix(fullKey, key) {
			return (link + opts), true
		}
	}

	// Search as suffix
	noSlashKey := key[1:]
	for fullKey, link := range config.LinkMap {
		if strings.HasSuffix(fullKey, noSlashKey) {
			return (link + opts), true
		}
	}

	return "", false
}
