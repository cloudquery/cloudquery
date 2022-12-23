package client

import "fmt"

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
	AdminAccount                *Account `json:"admin_account"`
	MemberCredentials           *Account `json:"member_trusted_principal"`
	ChildAccountRoleName        string   `json:"member_role_name,omitempty"`
	ChildAccountRoleSessionName string   `json:"member_role_session_name,omitempty"`
	ChildAccountExternalID      string   `json:"member_external_id,omitempty"`
	ChildAccountRegions         []string `json:"member_regions,omitempty"`
}

type Spec struct {
	Regions           []string  `json:"regions,omitempty"`
	Accounts          []Account `json:"accounts"`
	Organization      *AwsOrg   `json:"org"`
	AWSDebug          bool      `json:"aws_debug,omitempty"`
	MaxRetries        *int      `json:"max_retries,omitempty"`
	MaxBackoff        *int      `json:"max_backoff,omitempty"`
	EndpointURL       string    `json:"custom_endpoint_url,omitempty"`
	HostnameImmutable *bool     `json:"custom_endpoint_hostname_immutable,omitempty"`
	PartitionID       string    `json:"custom_endpoint_partition_id,omitempty"`
	SigningRegion     string    `json:"custom_endpoint_signing_region,omitempty"`
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
	return nil
}
