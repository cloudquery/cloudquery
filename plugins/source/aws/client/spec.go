package client

type Account struct {
	Name           string   `json:"name,omitempty"`
	LocalProfile          string   `json:"local_profile,omitempty"`
	AssumeRoleARN         string   `json:"assume_role_arn,omitempty"`
	AssumeRoleSessionName string   `json:"assume_role_session_name,omitempty"`
	AssumeRoleExternalID  string   `json:"assume_role_external_id,omitempty"`
	DefaultRegion         string   `json:"default_region,omitempty"`
	Regions               []string `json:"regions,omitempty"`
	source                string
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
	Regions      []string  `json:"regions,omitempty"`
	Accounts     []Account `json:"accounts,omitempty"`
	Organization *AwsOrg   `json:"org,omitempty"`
	Debug        bool      `json:"debug,omitempty"`
	MaxAttempts  int       `json:"max_attempts,omitempty"`
	MaxBackoff   int       `json:"max_backoff,omitempty"`
}

func (s *Spec) SetDefault() {
	if s.MaxAttempts == 0 {
		s.MaxAttempts = 10
	}
	if s.MaxBackoff == 0 {
		s.MaxBackoff = 20
	}
}
