package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/smithy-go"
)

//log-group:([a-zA-Z0-9/]+):
var GroupNameRegex = regexp.MustCompile("arn:aws:logs:[a-z0-9-]+:[0-9]+:log-group:([a-zA-Z0-9-/]+):")

type SupportedServicesData struct {
	Services map[string]struct {
		Regions []string `json:"regions"`
		Id      string   `json:"id"`
		Name    string   `json:"name"`
	} `json:"services"`
}

// supportedServices map of the supported service-regions
var supportedServices map[string]map[string]struct{}
var getSupportedServices sync.Once

// apiErrorServiceNames stores api subdomains and service names for error decoding
// some services have a few subdomains that differs from service name
// The list is not full todo take this list using some api
var apiErrorServiceNames = map[string]string{
	"mq":               "amazon-mq",
	"cognito-identity": "cognito",
	"cognito-idp":      "cognito",
	"acm":              "certificate-manager",
	"acm-pca":          "certificate-manager",
	"ce":               "aws-cost-management-cost-explorer",
	"groundstation":    "ground-station",
}

const supportedServicesLink = "https://raw.githubusercontent.com/burib/aws-region-table-parser/master/data/parseddata.json"

// downloadSupportedResourcesForRegions gets the data about AWS services and regions they are available in
func downloadSupportedResourcesForRegions() (map[string]map[string]struct{}, error) {
	req, err := http.NewRequest(http.MethodGet, supportedServicesLink, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get aws supported resources for region, status code: %d", resp.StatusCode)
	}

	var data SupportedServicesData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	m := make(map[string]map[string]struct{})
	for k, s := range data.Services {
		if _, ok := m[k]; !ok {
			m[k] = make(map[string]struct{})
		}
		for _, r := range s.Regions {
			m[k][r] = struct{}{}
		}
	}

	return m, nil
}

// ignoreUnsupportedResourceForRegionError returns true if request was sent to a service that exists but fetched with error
func ignoreUnsupportedResourceForRegionError(err error) bool {
	getSupportedServices.Do(func() {
		supportedServices, _ = downloadSupportedResourcesForRegions()
	})
	var dnsErr *net.DNSError
	if supportedServices != nil && errors.As(err, &dnsErr) && dnsErr.IsNotFound {
		var parts = make([]string, 0)

		// if error address in local, try to parse the original error
		if strings.HasPrefix(dnsErr.Name, "127.0.0.53:53") {
			var re = regexp.MustCompile(`(?m).*Post "https:\/\/(.*)?\.(.*)?\.amazonaws.com\/": dial tcp:.*`)
			if re.MatchString(err.Error()) {
				p := re.FindAllStringSubmatch(err.Error(), -1)
				if len(p) > 0 && len(p[0]) > 0 {
					parts = append(parts, p[0][1])
					parts = append(parts, p[0][2])
				}
			}
		} else {
			parts = strings.Split(dnsErr.Name, ".")
		}

		if len(parts) < 2 {
			// usual aws domain has more than 2 parts
			return false
		}
		apiService, ok := apiErrorServiceNames[parts[0]]
		if !ok {
			apiService = parts[0]
		}
		region := parts[1]

		_, ok = supportedServices[apiService][region]
		// if service-region combination is in the map than service is supported and error should not be ignored
		return !ok
	}
	return true
}

func IgnoreAccessDeniedServiceDisabled(err error) bool {
	if ignoreUnsupportedResourceForRegionError(err) {
		return true
	}
	var ae smithy.APIError
	if errors.As(err, &ae) {
		switch ae.ErrorCode() {
		case "AWSOrganizationsNotInUseException":
			return true
		case "AuthorizationError", "AccessDenied", "AccessDeniedException", "UnauthorizedOperation":
			return true
		case "OptInRequired", "SubscriptionRequiredException", "InvalidClientTokenId":
			return true
		}
	}
	return false
}

func IgnoreWithInvalidAction(err error) bool {
	if IgnoreAccessDeniedServiceDisabled(err) {
		return true
	}
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "InvalidAction" {
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

	// if resource type is empty
	// for example s3 bucket
	resource := ""
	if resourceType == "" {
		resource = resourceID
	} else {
		resource = fmt.Sprintf("%s/%s", resourceType, resourceID)
	}

	return arn.ARN{
		// TODO: Make this configurable in the future
		Partition: "aws",
		Service:   service,
		Region:    region,
		AccountID: accountID,
		Resource:  resource,
	}.String()
}

func accountObfusactor(aa []Account, msg string) string {
	for _, a := range aa {
		msg = strings.ReplaceAll(msg, a.ID, obfuscateAccountId(a.ID))
	}
	return msg
}
