package client

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

type AWSService string

type AwsService struct {
	Regions map[string]*map[string]any `json:"regions"`
}

type AwsPartition struct {
	Id       string                 `json:"partition"`
	Name     string                 `json:"partitionName"`
	Services map[string]*AwsService `json:"services"`
}

type SupportedServiceRegionsData struct {
	Partitions        map[string]AwsPartition `json:"partitions"`
	regionVsPartition map[string]string
}

// ListResolver is responsible for iterating through entire list of resources that should be grabbed (if API is paginated). It should send list of items via the `resultsChan` so that the DetailResolver can grab the details of each item. All errors should be sent to the error channel.
type ListResolverFunc func(ctx context.Context, meta schema.ClientMeta, detailChan chan<- any) error

// DetailResolveFunc is responsible for grabbing any and all metadata for a resource. All errors should be sent to the error channel.
type DetailResolverFunc func(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- any, errorChan chan<- error, summary any)

const (
	ApigatewayService           AWSService = "apigateway"
	Athena                      AWSService = "athena"
	CloudformationService       AWSService = "cloudformation"
	CloudfrontService           AWSService = "cloudfront"
	CognitoIdentityService      AWSService = "cognito-identity"
	DirectConnectService        AWSService = "directconnect"
	DynamoDBService             AWSService = "dynamodb"
	EC2Service                  AWSService = "ec2"
	EFSService                  AWSService = "elasticfilesystem"
	ElasticLoadBalancingService AWSService = "elasticloadbalancing"
	GlueService                 AWSService = "glue"
	GuardDutyService            AWSService = "guardduty"
	IamService                  AWSService = "iam"
	RedshiftService             AWSService = "redshift"
	Route53Service              AWSService = "route53"
	S3Service                   AWSService = "s3"
	SESService                  AWSService = "ses"
	WAFRegional                 AWSService = "waf-regional"
	WorkspacesService           AWSService = "workspaces"
	XRayService                 AWSService = "xray"
)

const (
	PartitionServiceRegionFile = "data/partition_service_region.json"
	defaultPartition           = "aws"
)

var (
	//go:embed data/partition_service_region.json
	supportedServiceRegionFile embed.FS
	readOnce                   sync.Once
	supportedServiceRegion     *SupportedServiceRegionsData
)

var notFoundErrorSubstrings = []string{
	"InvalidAMIID.Unavailable",
	"NonExistentQueue",
	"NoSuch",
	"NotFound",
	"ResourceNotFoundException",
	"WAFNonexistentItemException",
	"NoSuchResource",
}

var accessDeniedErrorStrings = map[string]struct{}{
	"AuthorizationError":              {},
	"AccessDenied":                    {},
	"AccessDeniedException":           {},
	"InsufficientPrivilegesException": {},
	"UnauthorizedOperation":           {},
	"Unauthorized":                    {},
}

func ReadSupportedServiceRegions() *SupportedServiceRegionsData {
	f, err := supportedServiceRegionFile.Open(PartitionServiceRegionFile)
	if err != nil {
		return nil
	}
	stat, err := f.Stat()
	if err != nil {
		return nil
	}
	data := make([]byte, stat.Size())
	if _, err := f.Read(data); err != nil {
		return nil
	}

	var result SupportedServiceRegionsData
	if err := json.Unmarshal(data, &result); err != nil {
		return nil
	}

	result.regionVsPartition = make(map[string]string)
	for _, p := range result.Partitions {
		for _, svc := range p.Services {
			for reg := range svc.Regions {
				result.regionVsPartition[reg] = p.Id
			}
		}
	}

	return &result
}

func supportedRegions(service string) []string {
	readOnce.Do(func() {
		supportedServiceRegion = ReadSupportedServiceRegions()
	})

	if supportedServiceRegion == nil {
		return nil
	}

	if supportedServiceRegion.Partitions == nil {
		return nil
	}
	regions := make([]string, 0)
	for id := range supportedServiceRegion.Partitions {
		currentPartition := supportedServiceRegion.Partitions[id]

		if currentPartition.Services[service] == nil {
			continue
		}

		for region := range currentPartition.Services[service].Regions {
			regions = append(regions, region)
		}
	}

	return regions
}

func isSupportedServiceForRegion(service string, region string) bool {
	for _, r := range supportedRegions(service) {
		if r == region {
			return true
		}
	}
	return false
}

func getAvailableRegions() (map[string]bool, error) {
	readOnce.Do(func() {
		supportedServiceRegion = ReadSupportedServiceRegions()
	})

	regionsSet := make(map[string]bool)

	if supportedServiceRegion == nil {
		return nil, fmt.Errorf("could not get AWS regions/services data")
	}

	if supportedServiceRegion.Partitions == nil {
		return nil, fmt.Errorf("could not found any AWS partitions")
	}

	for _, prt := range supportedServiceRegion.Partitions {
		for _, service := range prt.Services {
			for region := range service.Regions {
				regionsSet[region] = true
			}
		}
	}

	return regionsSet, nil
}

func RegionsPartition(region string) (string, bool) {
	readOnce.Do(func() {
		supportedServiceRegion = ReadSupportedServiceRegions()
	})

	prt, ok := supportedServiceRegion.regionVsPartition[region]
	if !ok {
		return defaultPartition, false
	}
	return prt, true
}

func IgnoreAccessDeniedServiceDisabled(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		switch ae.ErrorCode() {
		case "UnrecognizedClientException":
			return strings.Contains(ae.Error(), "The security token included in the request is invalid")
		case "AWSOrganizationsNotInUseException":
			return true
		case "OptInRequired", "SubscriptionRequiredException", "InvalidClientTokenId":
			return true
		}
	}
	return isAccessDeniedError(err)
}

func IgnoreCommonErrors(err error) bool {
	if IgnoreAccessDeniedServiceDisabled(err) || IgnoreNotAvailableRegion(err) || IgnoreWithInvalidAction(err) || isNotFoundError(err) {
		return true
	}
	return false
}

func IgnoreWithInvalidAction(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "InvalidAction" {
			return true
		}
	}
	return false
}

func IgnoreNotAvailableRegion(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "InvalidRequestException" && strings.Contains(ae.ErrorMessage(), "not available in the current Region") {
			return true
		}
	}
	return false
}

func resolveARN(service AWSService, resourceID func(resource *schema.Resource) ([]string, error), useRegion, useAccountID bool) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		cl := meta.(*Client)
		idParts, err := resourceID(resource)
		if err != nil {
			return fmt.Errorf("error resolving resource id: %w", err)
		}
		var accountID, region string
		if useAccountID {
			accountID = cl.AccountID
		}
		if useRegion {
			region = cl.Region
		}
		return resource.Set(c.Name, arn.ARN{
			Partition: cl.Partition,
			Service:   string(service),
			Region:    region,
			AccountID: accountID,
			Resource:  strings.Join(idParts, "/"),
		}.String())
	}
}

// ResolveARNWithAccount returns a column resolver that will set a field value to a proper ARN
// based on provided AWS service and resource id value returned by resourceID function.
// Region is left empty and account id is set to the value of the client.
func ResolveARNWithAccount(service AWSService, resourceID func(resource *schema.Resource) ([]string, error)) schema.ColumnResolver {
	return resolveARN(service, resourceID, false, true)
}

// ResolveARN returns a column resolver that will set a field value to a proper ARN
// based on provided AWS service and resource id value returned by resourceID function.
// Region and account id are set to the values of the client.
func ResolveARN(service AWSService, resourceID func(resource *schema.Resource) ([]string, error)) schema.ColumnResolver {
	return resolveARN(service, resourceID, true, true)
}

// ResolveARNGlobal returns a column resolver that will set a field value to a proper ARN
// based on provided AWS service and resource id value returned by resourceID function.
// Region  and account id are left empty.
func ResolveARNGlobal(service AWSService, resourceID func(resource *schema.Resource) ([]string, error)) schema.ColumnResolver {
	return resolveARN(service, resourceID, false, false)
}

// IsNotFoundError checks if api error should be ignored
func (c *Client) IsNotFoundError(err error) bool {
	if isNotFoundError(err) {
		c.logger.Warn().Err(err).Msg("API returned \"NotFound\" error ignoring it...")
		return true
	}
	return false
}

func isNotFoundError(err error) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	errorCode := ae.ErrorCode()
	for _, s := range notFoundErrorSubstrings {
		if strings.Contains(errorCode, s) {
			return true
		}
	}
	return false
}

func isAccessDeniedError(err error) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	_, ok := accessDeniedErrorStrings[ae.ErrorCode()]
	return ok
}

func IsInvalidParameterValueError(err error) bool {
	var apiErr smithy.APIError
	return errors.As(err, &apiErr) && apiErr.ErrorCode() == "InvalidParameterValue"
}

func IsAWSError(err error, code ...string) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	for _, c := range code {
		if strings.Contains(ae.ErrorCode(), c) {
			return true
		}
	}
	return false
}

// TagsIntoMap expects []T (usually "[]Tag") where T has "Key" and "Value" fields (of type string or *string) and writes them into the given map
func TagsIntoMap(tagSlice any, dst map[string]string) {
	stringify := func(v reflect.Value) string {
		vt := v.Type()
		if vt.Kind() == reflect.String {
			return v.String()
		}
		if vt.Kind() != reflect.Ptr || vt.Elem().Kind() != reflect.String {
			panic("field is not string or *string")
		}

		if v.IsNil() {
			// return empty string if string pointer is nil
			return ""
		}

		return v.Elem().String()
	}

	if k := reflect.TypeOf(tagSlice).Kind(); k != reflect.Slice {
		panic("invalid usage: Only slices are supported as input: " + k.String())
	}
	slc := reflect.ValueOf(tagSlice)

	for i := 0; i < slc.Len(); i++ {
		val := slc.Index(i)
		if k := val.Kind(); k != reflect.Struct {
			panic("slice member is not struct: " + k.String())
		}

		// key cannot be nil, but value can in the case of key-only tags
		keyField, valField := val.FieldByName("Key"), val.FieldByName("Value")
		if keyField.Type().Kind() == reflect.Ptr && keyField.IsNil() {
			continue
		}

		if keyField.IsZero() {
			panic("slice member is missing Key field")
		}

		dst[stringify(keyField)] = stringify(valField)
	}
}

// TagsToMap expects []T (usually "[]Tag") where T has "Key" and "Value" fields (of type string or *string) and returns a map
func TagsToMap(tagSlice any) map[string]string {
	if k := reflect.TypeOf(tagSlice).Kind(); k != reflect.Slice {
		panic("invalid usage: Only slices are supported as input: " + k.String())
	}
	slc := reflect.ValueOf(tagSlice)

	ret := make(map[string]string, slc.Len())
	TagsIntoMap(tagSlice, ret)
	return ret
}

func Sleep(ctx context.Context, dur time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(dur):
		return nil
	}
}

func CreateTrimPrefixTransformer(prefixes ...string) func(field reflect.StructField) (string, error) {
	return func(field reflect.StructField) (string, error) {
		name, err := transformers.DefaultNameTransformer(field)
		if err != nil {
			return "", err
		}
		for _, v := range prefixes {
			if strings.HasPrefix(name, v) {
				return name[len(v):], nil
			}
		}
		return name, nil
	}
}

func CreateReplaceTransformer(replace map[string]string) func(field reflect.StructField) (string, error) {
	return func(field reflect.StructField) (string, error) {
		name, err := transformers.DefaultNameTransformer(field)
		if err != nil {
			return "", err
		}
		for k, v := range replace {
			name = strings.ReplaceAll(name, k, v)
		}
		return name, nil
	}
}
