package client

import (
	"fmt"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"regexp"
	"strings"
)

const resourceIDPatternText = `(?i)subscriptions/(.+)/resourceGroups/(.+)/providers/(.+?)/(.+?)/(.+)`

var resourceIDPattern = regexp.MustCompile(resourceIDPatternText)

func DeleteSubscriptionFilter(meta schema.ClientMeta) []interface{} {
	client := meta.(*Client)
	return []interface{}{"subscription_id", client.SubscriptionId}
}

// ResourceDetails contains details about an Azure resource
type ResourceDetails struct {
	Subscription  string
	ResourceGroup string
	Provider      string
	ResourceType  string
	ResourceName  string
}

// ParseResourceID parses a resource ID into a ResourceDetails struct
func ParseResourceID(resourceID string) (ResourceDetails, error) {
	match := resourceIDPattern.FindStringSubmatch(resourceID)

	if len(match) == 0 {
		return ResourceDetails{}, fmt.Errorf("parsing failed for %s. Invalid resource Id format", resourceID)
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

	return result, nil
}
