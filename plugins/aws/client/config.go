package client

type Account struct {
	ID              string `hcl:",label"`
	AccountID       string
	AccountName     string   `hcl:"account_name,optional"`
	LocalProfile    string   `hcl:"local_profile,optional"`
	RoleARN         string   `hcl:"role_arn,optional"`
	RoleSessionName string   `hcl:"role_session_name,optional"`
	ExternalID      string   `hcl:"external_id,optional"`
	Regions         []string `hcl:"regions,optional"`
	source          string
}

type AwsOrg struct {
	OrganizationUnits           []string `hcl:"organization_units,optional"`
	AdminAccount                *Account `hcl:"admin_account,block"`
	MemberCredentials           *Account `hcl:"member_trusted_principal,block"`
	ChildAccountRoleName        string   `hcl:"member_role_name,optional"`
	ChildAccountRoleSessionName string   `hcl:"member_role_session_name,optional"`
	ChildAccountExternalID      string   `hcl:"member_external_id,optional"`
	ChildAccountRegions         []string `hcl:"member_regions,optional"`
}

type Config struct {
	Regions      []string  `hcl:"regions,optional"`
	Accounts     []Account `hcl:"accounts,block"`
	Organization *AwsOrg   `hcl:"org,block"`
	AWSDebug     bool      `hcl:"aws_debug,optional"`
	MaxRetries   int       `hcl:"max_retries,optional" default:"10"`
	MaxBackoff   int       `hcl:"max_backoff,optional" default:"30"`
}

func (c Config) Example() string {
	return ` configuration {
  // Optional, Repeated. Add an 'accounts' block for every account you want to assume-role into and fetch data from.
  // accounts "<UNIQUE ACCOUNT IDENTIFIER>" {
    // Optional. Role ARN we want to assume when accessing this account
    // role_arn = < YOUR_ROLE_ARN >
    // Optional. Named profile in config or credential file from where CQ should grab credentials
    // local_profile = < PROFILE_NAME >
  // }
  // Optional. by default assumes all regions
  // regions = ["us-east-1", "us-west-2"]
  // Optional. Enable AWS SDK debug logging.
  aws_debug = false
  // The maximum number of times that a request will be retried for failures. Defaults to 10 retry attempts.
  // max_retries = 10
  // The maximum back off delay between attempts. The backoff delays exponentially with a jitter based on the number of attempts. Defaults to 30 seconds.
  // max_backoff = 30
}
`
}
