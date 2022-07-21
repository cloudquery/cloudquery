package client

type Account struct {
	ID              string `yaml:"id"`
	AccountID       string
	AccountName     string   `yaml:"account_name,omitempty"`
	LocalProfile    string   `yaml:"local_profile,omitempty"`
	RoleARN         string   `yaml:"role_arn,omitempty"`
	RoleSessionName string   `yaml:"role_session_name,omitempty"`
	ExternalID      string   `yaml:"external_id,omitempty"`
	DefaultRegion   string   `yaml:"default_region,omitempty"`
	Regions         []string `yaml:"regions,omitempty"`
	source          string
}

type AwsOrg struct {
	OrganizationUnits           []string `yaml:"organization_units,omitempty"`
	AdminAccount                *Account `yaml:"admin_account"`
	MemberCredentials           *Account `yaml:"member_trusted_principal"`
	ChildAccountRoleName        string   `yaml:"member_role_name,omitempty"`
	ChildAccountRoleSessionName string   `yaml:"member_role_session_name,omitempty"`
	ChildAccountExternalID      string   `yaml:"member_external_id,omitempty"`
	ChildAccountRegions         []string `yaml:"member_regions,omitempty"`
}

type Config struct {
	Regions      []string  `yaml:"regions,omitempty"`
	Accounts     []Account `yaml:"accounts"`
	Organization *AwsOrg   `yaml:"org"`
	AWSDebug     bool      `yaml:"aws_debug,omitempty"`
	MaxRetries   int       `yaml:"max_retries,omitempty" default:"10"`
	MaxBackoff   int       `yaml:"max_backoff,omitempty" default:"30"`
	GlobalRegion string    `yaml:"global_region,omitempty" default:"us-east-1"`
}

func (Config) Example() string {
	return `
Optional, Repeated. Add an accounts block for every account you want to assume-role into and fetch data from.
accounts:
  - id: <UNIQUE ACCOUNT IDENTIFIER>
Optional. Role ARN we want to assume when accessing this account
    role_arn: < YOUR_ROLE_ARN >
Optional. Named profile in config or credential file from where CQ should grab credentials
    local_profile: < PROFILE_NAME >
Optional. by default assumes all regions
regions:
  - us-east-1
  - us-west-2
Optional. Enable AWS SDK debug logging.
  aws_debug: false
The maximum number of times that a request will be retried for failures. Defaults to 10 retry attempts.
max_retries: 10
The maximum back off delay between attempts. The backoff delays exponentially with a jitter based on the number of attempts. Defaults to 30 seconds.
max_backoff: 30
`
}
