package spec

import (
	"fmt"
	"regexp"
)

// Organization mode spec used to source all accounts underneath automatically.
type Organization struct {
	// Configuration for how to grab credentials from an admin account.
	AdminAccount *Account `json:"admin_account"`

	// Configuration for how to specify the principle to use in order to assume a role in the member accounts.
	MemberCredentials *Account `json:"member_trusted_principal"`

	// Role name that CloudQuery should use to assume a role in the member account from the admin account.
	//
	// Note: This is not a full ARN, it is just the name.
	ChildAccountRoleName string `json:"member_role_name,omitempty" jsonschema:"required,minLength=1,example=some_role_name"`

	// Overrides the default session name.
	ChildAccountRoleSessionName string `json:"member_role_session_name,omitempty" jsonschema:"example=some_role_session_name"`

	// Specify an external ID for use in the trust policy.
	ChildAccountExternalID string `json:"member_external_id,omitempty" jsonschema:"example=external_id"`

	// Limit fetching resources within this specific account to only these regions.
	// This will override any regions specified in the provider block.
	// You can specify all regions by using the `*` character as the only argument in the array.
	ChildAccountRegions []string `json:"member_regions,omitempty" jsonschema:"minLength=1,example=us-east-1"`

	// List of Organizational Units that CloudQuery should use to source accounts from.
	// If you specify an OU, CloudQuery will also traverse nested OUs.
	OrganizationUnits []string `json:"organization_units,omitempty" jsonschema:"pattern=^((ou-[0-9a-z]{4\\,32}-[a-z0-9]{8\\,32})|(r-[0-9a-z]{4\\,32}))$,example=ou-1234-12345678"`

	// List of Organizational Units to skip.
	// This is useful in conjunction with `organization_units` if there are child OUs that should be ignored.
	SkipOrganizationalUnits []string `json:"skip_organization_units,omitempty" jsonschema:"pattern=^((ou-[0-9a-z]{4\\,32}-[a-z0-9]{8\\,32})|(r-[0-9a-z]{4\\,32}))$,example=ou-1234-12345678"`

	// List of OU member accounts to skip.
	// This is useful if there are accounts under the selected OUs that should be ignored.
	SkipMemberAccounts []string `json:"skip_member_accounts,omitempty" jsonschema:"example=my_aws_account"`
}

func (o *Organization) Validate() error {
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
