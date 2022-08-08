package client

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
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

func DeleteSubscriptionFilter(meta schema.ClientMeta, _ *schema.Resource) []interface{} {
	client := meta.(*Client)
	return []interface{}{"subscription_id", client.SubscriptionId}
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

// ScopeSubscription returns a scope for the given subscription
func ScopeSubscription(subscriptionID string) string {
	return "subscriptions/" + subscriptionID
}

func IgnoreAccessDenied(err error) bool {
	var detailedError autorest.DetailedError

	return errors.As(err, &detailedError) && detailedError.StatusCode == http.StatusForbidden
}
