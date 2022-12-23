package client

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/thoas/go-funk"
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

	if s.Organization != nil && len(s.Accounts) > 0 {
		return errors.New("specifying accounts via both the Accounts and Org properties is not supported. To achieve both, use multiple source configurations")
	}
	if s.Organization != nil {
		if s.Organization.ChildAccountRoleName == "" {
			return fmt.Errorf("member_role_name is required when using org configuration")
		}
		if err := validateAccounts(s.Organization.SkipMemberAccounts); err != nil {
			return fmt.Errorf("invalid skip_member_accounts: %w", err)
		}
		if err := validateOUs(s.Organization.OrganizationUnits); err != nil {
			return fmt.Errorf("invalid organization_units: %w", err)
		}
		if err := validateOUs(s.Organization.SkipOrganizationalUnits); err != nil {
			return fmt.Errorf("invalid skip_organization_units: %w", err)
		}
	}
	if s.Accounts != nil {
		if err := validateAccounts(funk.Get(s.Accounts, "ID").([]string)); err != nil {
			return fmt.Errorf("invalid accounts: %w", err)
		}
	}
	return nil
}

func validateAccounts(accounts []string) error {
	r := regexp.MustCompile(`^(\d{12})$`)
	for _, account := range accounts {
		if !r.MatchString(account) {
			return fmt.Errorf("invalid account id: %s (should be 12 digits)", account)
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
