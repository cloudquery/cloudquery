package spec

import (
	"fmt"
	"regexp"
)

type Org struct {
	AdminAccount                *Account `json:"admin_account"`
	MemberCredentials           *Account `json:"member_trusted_principal"`
	ChildAccountRoleName        string   `json:"member_role_name,omitempty" jsonschema:"required,minLength=1"`
	ChildAccountRoleSessionName string   `json:"member_role_session_name,omitempty"`
	ChildAccountExternalID      string   `json:"member_external_id,omitempty"`
	ChildAccountRegions         []string `json:"member_regions,omitempty"`
	OrganizationUnits           []string `json:"organization_units,omitempty" jsonschema:"pattern=^((ou-[0-9a-z]{4\\,32}-[a-z0-9]{8\\,32})|(r-[0-9a-z]{4\\,32}))$"`
	SkipOrganizationalUnits     []string `json:"skip_organization_units,omitempty" jsonschema:"pattern=^((ou-[0-9a-z]{4\\,32}-[a-z0-9]{8\\,32})|(r-[0-9a-z]{4\\,32}))$"`
	SkipMemberAccounts          []string `json:"skip_member_accounts,omitempty"`
}

func (o *Org) Validate() error {
	if o.ChildAccountRoleName == "" {
		return fmt.Errorf("member_role_name is required when using org configuration")
	}
	if err := validateOUs(o.OrganizationUnits); err != nil {
		return fmt.Errorf("invalid organization_units: %w", err)
	}
	if err := validateOUs(o.SkipOrganizationalUnits); err != nil {
		return fmt.Errorf("invalid skip_organization_units: %w", err)
	}
	return nil
}

func validateOUs(ous []string) error {
	r := regexp.MustCompile(`^((ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})|(r-[0-9a-z]{4,32}))$`)
	for _, ou := range ous {
		if !r.MatchString(ou) {
			return fmt.Errorf(`invalid OU: %s (should match "ou-*-*" or "r-*" with lowercase letters or digits)`, ou)
		}
	}
	return nil
}
