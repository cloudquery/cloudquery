package client

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/smithy-go"
)

//log-group:([a-zA-Z0-9/]+):
var GroupNameRegex = regexp.MustCompile("arn:aws:logs:[a-z0-9-]+:[0-9]+:log-group:([a-zA-Z0-9-/]+):")

func IgnoreAccessDeniedServiceDisabled(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		switch ae.ErrorCode() {
		case "AccessDenied", "AccessDeniedException", "UnauthorizedOperation":
			return true
		case "OptInRequired", "SubscriptionRequiredException", "InvalidClientTokenId":
			return true
		}
	}
	return false
}

// GenerateResourceARN generates the arn for a resource.
// Service: The service name e.g. waf or elb or s3
// ResourceType: The sub resource type e.g. rule or instance (for an ec2 instance)
// ResourceID: The resource id e.g. i-1234567890abcdefg
// Region: The resource region e.g. us-east-1
// AccountID: The account id e.g. 123456789012
// See https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html for
// more information.
func GenerateResourceARN(service, resourceType, resourceID, region, accountID string) string {
	return arn.ARN{
		// TODO: Make this configurable in the future
		Partition: "aws",
		Service:   service,
		Region:    region,
		AccountID: accountID,
		Resource:  fmt.Sprintf("%s/%s", resourceType, resourceID),
	}.String()
}
