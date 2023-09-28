package spec

type Account struct {
	ID              string   `json:"id" jsonschema:"required,minLength=1"`
	AccountName     string   `json:"account_name,omitempty"`
	LocalProfile    string   `json:"local_profile,omitempty"`
	RoleARN         string   `json:"role_arn,omitempty" jsonschema:"pattern=^arn(:[^:\n]*){5\\,}([:/].*)?$"`
	RoleSessionName string   `json:"role_session_name,omitempty"`
	ExternalID      string   `json:"external_id,omitempty"`
	DefaultRegion   string   `json:"default_region,omitempty" jsonschema:"default=us-east-1"`
	Regions         []string `json:"regions,omitempty"`

	// explicitly ignore in JSON parsing, as this is filled in later
	Source AccountSource `json:"-"`
}

type AccountSource string

const (
	AccountSourceOrg = "org"
)
