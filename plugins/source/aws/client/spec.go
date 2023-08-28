package client

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
)

const (
	defaultMaxConcurrency = 50000
)

type Account struct {
	ID              string   `json:"id"`
	AccountName     string   `json:"account_name,omitempty"`
	LocalProfile    string   `json:"local_profile,omitempty"`
	RoleARN         string   `json:"role_arn,omitempty"`
	RoleSessionName string   `json:"role_session_name,omitempty"`
	ExternalID      string   `json:"external_id,omitempty"`
	DefaultRegion   string   `json:"default_region,omitempty"`
	Regions         []string `json:"regions,omitempty"`
	source          string
}

type AwsOrg struct {
	OrganizationUnits           []string `json:"organization_units,omitempty"`
	SkipMemberAccounts          []string `json:"skip_member_accounts,omitempty"`
	SkipOrganizationalUnits     []string `json:"skip_organization_units,omitempty"`
	AdminAccount                *Account `json:"admin_account"`
	MemberCredentials           *Account `json:"member_trusted_principal"`
	ChildAccountRoleName        string   `json:"member_role_name,omitempty"`
	ChildAccountRoleSessionName string   `json:"member_role_session_name,omitempty"`
	ChildAccountExternalID      string   `json:"member_external_id,omitempty"`
	ChildAccountRegions         []string `json:"member_regions,omitempty"`
}

type StreamingSync struct {
	Account           Account    `json:"account"`
	KinesisStreamARN  string     `json:"kinesis_stream_arn"`
	StartTime         *time.Time `json:"start_time,omitempty"`
	StreamingSyncOnly bool       `json:"streaming_sync_only"`
}
type Spec struct {
	Regions                   []string                   `json:"regions,omitempty"`
	Accounts                  []Account                  `json:"accounts"`
	Organization              *AwsOrg                    `json:"org"`
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
	StreamingSync             []StreamingSync            `json:"streaming_sync,omitempty"`
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
		if s.Organization.ChildAccountRoleName == "" {
			return fmt.Errorf("member_role_name is required when using org configuration")
		}
		if err := validateOUs(s.Organization.OrganizationUnits); err != nil {
			return fmt.Errorf("invalid organization_units: %w", err)
		}
		if err := validateOUs(s.Organization.SkipOrganizationalUnits); err != nil {
			return fmt.Errorf("invalid skip_organization_units: %w", err)
		}
	}
	if s.TableOptions != nil {
		if err := s.TableOptions.Validate(); err != nil {
			return fmt.Errorf("invalid table_options: %w", err)
		}
	}

	if len(s.StreamingSync) > 1 {
		return fmt.Errorf("only one streaming_sync is allowed at this time")
	}
	if len(s.StreamingSync) == 1 {
		_, err := arn.Parse(s.StreamingSync[0].KinesisStreamARN)
		if err != nil {
			return fmt.Errorf("failed to parse kinesis arn (%s): %w", s.StreamingSync[0].KinesisStreamARN, err)
		}
	}
	return nil
}

func validateOUs(ous []string) error {
	r := regexp.MustCompile(`^((ou\-[0-9a-z]{4,32}\-[a-z0-9]{8,32})|(r\-[0-9a-z]{4,32}))$`)
	for _, ou := range ous {
		if !r.MatchString(ou) {
			return fmt.Errorf(`invalid OU: %s (should match "ou-*-*" or "r-*" with lowercase letters or digits)`, ou)
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
}
