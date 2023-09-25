package spec

type Account struct {
	ID              string   `json:"id"`
	AccountName     string   `json:"account_name,omitempty"`
	LocalProfile    string   `json:"local_profile,omitempty"`
	RoleARN         string   `json:"role_arn,omitempty"`
	RoleSessionName string   `json:"role_session_name,omitempty"`
	ExternalID      string   `json:"external_id,omitempty"`
	DefaultRegion   string   `json:"default_region,omitempty"`
	Regions         []string `json:"regions,omitempty"`

	// explicitly ignore in JSON parsing, as this is filled in later
	Source AccountSource `json:"-"`
}

type AccountSource string

const (
	AccountSourceOrg = "org"
)
