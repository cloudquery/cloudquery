package spec

import (
	_ "embed"
	"errors"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
)

const (
	defaultMaxConcurrency = 50000
)

type Spec struct {
	Regions                   []string                   `json:"regions,omitempty"`
	Accounts                  []Account                  `json:"accounts"`
	Organization              *Org                       `json:"org"`
	AWSDebug                  bool                       `json:"aws_debug,omitempty"`
	MaxRetries                *int                       `json:"max_retries,omitempty"`
	MaxBackoff                *int                       `json:"max_backoff,omitempty"`
	EndpointURL               string                     `json:"custom_endpoint_url,omitempty"`
	HostnameImmutable         *bool                      `json:"custom_endpoint_hostname_immutable,omitempty"`
	PartitionID               string                     `json:"custom_endpoint_partition_id,omitempty"`
	SigningRegion             string                     `json:"custom_endpoint_signing_region,omitempty"`
	InitializationConcurrency int                        `json:"initialization_concurrency"`
	UsePaidAPIs               bool                       `json:"use_paid_apis"`
	TableOptions              *tableoptions.TableOptions `json:"table_options,omitempty"`
	Concurrency               int                        `json:"concurrency"`
	EventBasedSync            *EventBasedSync            `json:"event_based_sync,omitempty"`
	Scheduler                 scheduler.Strategy         `json:"scheduler,omitempty"`
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
	if s.InitializationConcurrency <= 0 {
		s.InitializationConcurrency = 4
	}
	if s.TableOptions == nil {
		s.TableOptions = &tableoptions.TableOptions{}
	}
	if s.Concurrency == 0 {
		s.Concurrency = defaultMaxConcurrency
	}
	if s.EventBasedSync != nil && s.EventBasedSync.FullSync == nil {
		fullSync := true
		s.EventBasedSync.FullSync = &fullSync
	}
}

// JSONSchemaExtend is required to remove stale definitions`.
// We use value receiver because of https://github.com/invopop/jsonschema/issues/102
//func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
//}

//go:embed schema.json
var JSONSchema string
