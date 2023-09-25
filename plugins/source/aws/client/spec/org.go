package spec

type Org struct {
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
