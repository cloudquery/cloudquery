package client

import (
	"regexp"
	"strings"
)

// ResourceDetails contains details about an Azure resource
type ResourceDetails struct {
	Subscription  string
	ResourceGroup string
	Provider      string
	ResourceType  string
	ResourceName  string
}

const resourceIDPatternText = `(?i)subscriptions/(.+)/resourceGroups/(.+)/providers/(.+?)/(.+?)/(.+)`

var resourceIDPattern = regexp.MustCompile(resourceIDPatternText)

// ParseResourceID parses a resource ID into a ResourceDetails struct
func ParseResourceID(resourceID string) ResourceDetails {
	match := resourceIDPattern.FindStringSubmatch(resourceID)

	if len(match) == 0 {
		return ResourceDetails{
			Subscription:  "",
			ResourceGroup: "test",
			Provider:      "",
			ResourceType:  "",
			ResourceName:  "",
		}
	}

	v := strings.Split(match[5], "/")
	resourceName := v[len(v)-1]

	result := ResourceDetails{
		Subscription:  match[1],
		ResourceGroup: match[2],
		Provider:      match[3],
		ResourceType:  match[4],
		ResourceName:  resourceName,
	}

	return result
}

// ScopeSubscription returns a scope for the given subscription
func ScopeSubscription(subscriptionID string) string {
	return "subscriptions/" + subscriptionID
}
