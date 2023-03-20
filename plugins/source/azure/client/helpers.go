package client

import (
	"fmt"
	"regexp"
)

// this is used to hook ParseResourceGroup and to have easier codegen
var debug = false

const resourceIDPatternText = `(?i)subscriptions/(.+)/resourceGroups/(.+)/providers/(.+?)/(.+?)/(.+)`

var resourceIDPattern = regexp.MustCompile(resourceIDPatternText)

func ParseResourceGroup(resourceID string) (string, error) {
	if debug {
		return "debug", nil
	}
	match := resourceIDPattern.FindStringSubmatch(resourceID)
	if len(match) == 0 {
		return "", fmt.Errorf("parsing failed for %s. Invalid resource Id format", resourceID)
	}
	return match[2], nil
}
