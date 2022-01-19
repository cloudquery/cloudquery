package client

type Account struct {
	ID           string   `hcl:",label"`
	AccountID    string   `hcl:"account_id,optional"`
	AccountName  string   `hcl:"account_name,optional"`
	RoleARN      string   `hcl:"role_arn,optional"`
	LocalProfile string   `hcl:"local_profile,optional"`
	ExternalID   string   `hcl:"external_id,optional"`
	Regions      []string `hcl:"regions,optional"`
}

type Config struct {
	Regions    []string  `hcl:"regions,optional"`
	Accounts   []Account `hcl:"accounts,block"`
	AWSDebug   bool      `hcl:"aws_debug,optional"`
	MaxRetries int       `hcl:"max_retries,optional" default:"10"`
	MaxBackoff int       `hcl:"max_backoff,optional" default:"90"`
}

func (c Config) Example() string {
	return ` configuration {
  // Optional. if you want to assume role to multiple account and fetch data from them
  // accounts "<YOUR ACCOUNT ID>" {
    // Optional. Role ARN we want to assume when accessing this account
    // role_arn = < YOUR_ROLE_ARN >
    // Optional. Account ID we want to assume when accessing this account - override the block label
    // account_id = < YOUR ACCOUNT ID >
  // }
  // Optional. by default assumes all regions
  // regions = ["us-east-1", "us-west-2"]
  // Optional. Enable AWS SDK debug logging.
  aws_debug = false
  // The maximum number of times that a request will be retried for failures. Defaults to 5 retry attempts.
  // max_retries = 5
  // The maximum back off delay between attempts. The backoff delays exponentially with a jitter based on the number of attempts. Defaults to 60 seconds.
  // max_backoff = 30 
}
`
}
