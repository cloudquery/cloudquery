package spec

import (
	_ "embed"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type Spec struct {
	// Regions to use.
	Regions []string `json:"regions,omitempty" jsonschema:"minLength=1,example=us-east-1"`

	// List of all accounts to fetch information from.
	Accounts []Account `json:"accounts"`

	// In AWS organization mode, CloudQuery will source all accounts underneath automatically.
	Organization *Organization `json:"org"`

	// If `true`, will log AWS debug logs, including retries and other request/response metadata. Requires passing `--log-level debug` to the CloudQuery CLI.
	AWSDebug bool `json:"aws_debug,omitempty" jsonschema:"default=false"`

	// Defines the maximum number of times an API request will be retried.
	MaxRetries *int `json:"max_retries,omitempty" jsonschema:"default=10"`

	// Defines the duration between retry attempts.
	MaxBackoff *int `json:"max_backoff,omitempty" jsonschema:"default=30"`

	// The base URL endpoint the SDK API clients will use to make API calls to.
	// The SDK will suffix URI path and query elements to this endpoint.
	EndpointURL string `json:"custom_endpoint_url,omitempty"`

	// Specifies if the endpoint's hostname can be modified by the SDK's API client.
	// When using something like LocalStack make sure to set it equal to `true`.
	HostnameImmutable *bool `json:"custom_endpoint_hostname_immutable,omitempty" jsonschema:"default=false"`

	// The AWS partition the endpoint belongs to.
	PartitionID string `json:"custom_endpoint_partition_id,omitempty" jsonschema:"default=,example=aws"`

	// The region that should be used for signing the request to the endpoint.
	SigningRegion string `json:"custom_endpoint_signing_region,omitempty" jsonschema:"default=,example=us-east-1"`

	// During initialization the AWS source plugin fetches information about each account and region.
	// This setting controls how many accounts can be initialized concurrently.
	// Only configurations with many accounts (either hardcoded or discovered via Organizations)
	// should require modifying this setting, to either lower it to avoid rate limit errors, or to increase it to speed up the initialization process.
	InitializationConcurrency int `json:"initialization_concurrency" jsonschema:"minimum=1,default=4"`

	// The best effort maximum number of Go routines to use. Lower this number to reduce memory usage.
	Concurrency int `json:"concurrency" jsonschema:"minimum=1,default=50000"`

	// When set to `true` plugin will sync data from APIs that incur a fee.
	UsePaidAPIs bool `json:"use_paid_apis" jsonschema:"default=false"`

	// The scheduler to use when determining the priority of resources to sync. By default, it is set to `shuffle`.
	//
	// For more information about this, see [performance tuning](/docs/advanced-topics/performance-tuning).
	Scheduler *scheduler.Strategy `json:"scheduler,omitempty" jsonschema:"default=shuffle"`
}

// JSONSchemaExtend is required to verify:
// 1.if `custom_endpoint_url` is present then the following fields are required:
// * `custom_endpoint_partition_id`
// * `custom_endpoint_signing_region`
// * `custom_endpoint_hostname_immutable`
// 2. Make `org` & `accounts` mutually exclusive
func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.AllOf = []*jsonschema.Schema{
		{
			// custom_endpoint_url => custom_endpoint_partition_id, custom_endpoint_signing_region, custom_endpoint_hostname_immutable
			If: &jsonschema.Schema{
				// We also need to make sure that `custom_endpoint_url` isn't ""
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					properties := jsonschema.NewProperties()
					url := *sc.Properties.Value("custom_endpoint_url")
					url.MinLength = aws.Uint64(1)
					properties.Set("custom_endpoint_url", &url)
					return properties
				}(),
				Required: []string{"custom_endpoint_url"},
			},
			Then: &jsonschema.Schema{
				// require properties not to be empty or null
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					properties := jsonschema.NewProperties()

					partitionID := *sc.Properties.Value("custom_endpoint_partition_id")
					partitionID.MinLength = aws.Uint64(1)
					properties.Set("custom_endpoint_partition_id", &partitionID)

					signingRegion := *sc.Properties.Value("custom_endpoint_signing_region")
					signingRegion.MinLength = aws.Uint64(1)
					properties.Set("custom_endpoint_signing_region", &signingRegion)

					hostnameImmutable := *sc.Properties.Value("custom_endpoint_hostname_immutable").OneOf[0] // spec is 0, null is 1st
					properties.Set("custom_endpoint_hostname_immutable", &hostnameImmutable)

					return properties
				}(),
				Required: []string{"custom_endpoint_partition_id", "custom_endpoint_signing_region", "custom_endpoint_hostname_immutable"},
			},
		},
		{
			Not: &jsonschema.Schema{
				// org & accounts are mutually exclusive
				Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
					properties := jsonschema.NewProperties()
					properties.Set("org", sc.Properties.Value("org").OneOf[0]) // spec is 0, null is 1st

					// we take a value because we'll be updating the items spec
					accounts := *sc.Properties.Value("accounts").OneOf[0] // spec is 0, null is 1st
					accounts.MinItems = aws.Uint64(1)
					properties.Set("accounts", &accounts)
					return properties
				}(),
				Required: []string{"org", "accounts"},
			},
		},
	}

	// additionally, as we want to outline non-sdk-default for scheduler strategy
	sc.Properties.Value("scheduler").Default = (&[]scheduler.Strategy{scheduler.StrategyShuffle}[0]).String()
}

func (s *Spec) Validate() error {
	if s.EndpointURL != "" {
		if s.PartitionID == "" {
			return fmt.Errorf("custom_endpoint_partition_id is required when custom_endpoint_url is set")
		}
		if s.SigningRegion == "" {
			return fmt.Errorf("custom_endpoint_signing_region is required when custom_endpoint_url is set")
		}
		if s.HostnameImmutable == nil {
			return fmt.Errorf("custom_endpoint_hostname_immutable is required when custom_endpoint_url is set")
		}
	}

	if s.Organization != nil && len(s.Accounts) > 0 {
		return errors.New("specifying accounts via both the Accounts and Organization properties is not supported. To achieve both, use multiple source configurations")
	}
	if s.Organization != nil {
		if err := s.Organization.Validate(); err != nil {
			return fmt.Errorf("invalid org: %w", err)
		}
	}

	return nil
}

func (s *Spec) SetDefaults() {
	if s.InitializationConcurrency <= 0 {
		const defaultInitializationConcurrency = 4
		s.InitializationConcurrency = defaultInitializationConcurrency
	}

	if s.Concurrency <= 0 {
		const defaultMaxConcurrency = 50000
		s.Concurrency = defaultMaxConcurrency
	}

	if s.MaxRetries == nil {
		maxRetries := 10
		s.MaxRetries = &maxRetries
	}

	if s.MaxBackoff == nil {
		maxBackoff := 30
		s.MaxBackoff = &maxBackoff
	}

	if s.Scheduler == nil {
		strategy := scheduler.StrategyShuffle
		s.Scheduler = &strategy
	}
}

//go:embed schema.json
var JSONSchema string
