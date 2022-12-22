package client

import (
	"fmt"
	"strings"
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
	Regions      []string  `json:"regions,omitempty"`
	Accounts     []Account `json:"accounts"`
	Organization *AwsOrg   `json:"org"`
	AWSDebug     bool      `json:"aws_debug,omitempty"`
	MaxRetries   *int      `json:"max_retries,omitempty"`
	MaxBackoff   *int      `json:"max_backoff,omitempty"`
}

func (s *Spec) Validate() error {
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
}

func validateOUs(ous []string) error {
	for _, ou := range ous {
		if !strings.HasPrefix(ou, "ou-") {
			return fmt.Errorf("invalid OU: %s (should match ou-*)", ou)
		}
	}
}
