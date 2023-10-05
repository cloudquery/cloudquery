package spec

import (
	_ "embed"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/spec/tableoptions"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type Spec struct {
	Regions                   []string                   `json:"regions,omitempty" jsonschema:"minLength=1,default=us-east-1"`
	Accounts                  []Account                  `json:"accounts"`
	Organization              *Org                       `json:"org"`
	AWSDebug                  bool                       `json:"aws_debug,omitempty"`
	MaxRetries                *int                       `json:"max_retries,omitempty" jsonschema:"default=10"`
	MaxBackoff                *int                       `json:"max_backoff,omitempty" jsonschema:"default=30"`
	EndpointURL               string                     `json:"custom_endpoint_url,omitempty"`
	HostnameImmutable         *bool                      `json:"custom_endpoint_hostname_immutable,omitempty"`
	PartitionID               string                     `json:"custom_endpoint_partition_id,omitempty"`
	SigningRegion             string                     `json:"custom_endpoint_signing_region,omitempty"`
	InitializationConcurrency int                        `json:"initialization_concurrency" jsonschema:"minimum=1,default=4"`
	Concurrency               int                        `json:"concurrency" jsonschema:"minimum=1,default=50000"`
	UsePaidAPIs               bool                       `json:"use_paid_apis" jsonschema:"default=false"`
	TableOptions              *tableoptions.TableOptions `json:"table_options,omitempty"`
	EventBasedSync            *EventBasedSync            `json:"event_based_sync,omitempty"`
	Scheduler                 scheduler.Strategy         `json:"scheduler,omitempty"`
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
		return errors.New("specifying accounts via both the Accounts and Org properties is not supported. To achieve both, use multiple source configurations")
	}
	if s.Organization != nil {
		if err := s.Organization.Validate(); err != nil {
			return fmt.Errorf("invalid org: %w", err)
		}
	}

	if s.TableOptions != nil {
		if err := s.TableOptions.Validate(); err != nil {
			return fmt.Errorf("invalid table_options: %w", err)
		}
	}

	if s.EventBasedSync != nil {
		if err := s.EventBasedSync.Validate(); err != nil {
			return fmt.Errorf("invalid event_based_sync: %w", err)
		}
	}
	return nil
}

func (s *Spec) SetDefaults() {
	if s.TableOptions == nil {
		s.TableOptions = &tableoptions.TableOptions{}
	}
	// also call set defaults
	s.TableOptions.SetDefaults()

	if s.InitializationConcurrency <= 0 {
		const defaultInitializationConcurrency = 4
		s.InitializationConcurrency = defaultInitializationConcurrency
	}

	if s.Concurrency <= 0 {
		const defaultMaxConcurrency = 50000
		s.Concurrency = defaultMaxConcurrency
	}

	if s.EventBasedSync != nil && s.EventBasedSync.FullSync == nil {
		fullSync := true
		s.EventBasedSync.FullSync = &fullSync
	}

	if s.MaxRetries == nil {
		maxRetries := 10
		s.MaxRetries = &maxRetries
	}

	if s.MaxBackoff == nil {
		maxBackoff := 30
		s.MaxBackoff = &maxBackoff
	}
}

//go:embed schema.json
var JSONSchema string
