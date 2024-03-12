package spec

// This is used to specify one or more accounts to extract information from.
type Account struct {
	// Will be used as an alias in the source plugin and in the logs.
	ID string `json:"id" jsonschema:"required,minLength=1,example=some_id"`

	// Will be used as an alias in the source plugin and in the logs.
	AccountName string `json:"account_name,omitempty" jsonschema:"example=my_aws_account_name"`

	// [Local profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) to use to authenticate this account with.
	// Please note this should be set to the name of the profile.
	//
	// For example, with the following credentials file:
	//
	//   ```ini copy
	//   [default]
	//   aws_access_key_id=xxxx
	//   aws_secret_access_key=xxxx
	//
	//   [user1]
	//   aws_access_key_id=xxxx
	//   aws_secret_access_key=xxxx
	//   ```
	//
	// `local_profile` should be set to either `default` or `user1`.
	LocalProfile string `json:"local_profile,omitempty" jsonschema:"example=my_aws_profile"`

	// If specified will use this to assume role.
	RoleARN string `json:"role_arn,omitempty" jsonschema:"pattern=^(arn(:[^:\n]*){5}([:/].*)?)?$"`

	// If specified will use this session name when assume role to `role_arn`.
	RoleSessionName string `json:"role_session_name,omitempty" jsonschema:"example=my_aws_role_session_name"`

	// If specified will use this when assuming role to `role_arn`.
	ExternalID string `json:"external_id,omitempty" jsonschema:"example=external_id"`

	// If specified, this region will be used as the default region for the account.
	DefaultRegion string `json:"default_region,omitempty" jsonschema:"minLength=1,default=us-east-1"`

	// Regions to use for this account. Defaults to global `regions` setting.
	Regions []string `json:"regions,omitempty" jsonschema:"minLength=1,example=us-east-1"`

	// explicitly ignore in JSON parsing, as this is filled in later
	Source AccountSource `json:"-"`
}

type AccountSource string

const (
	AccountSourceOrg = "org"
)
