# AWS Source Plugin Configuration Reference

## Example

This example connects a single AWS account in one region to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](https://www.cloudquery.io/docs/reference/source-spec).

```yml
kind: source
spec:
  # Source spec section
  name: aws
  path: cloudquery/aws
  version: "v4.1.0" # latest version of aws plugin
  tables: ["*"]
  destinations: ["postgresql"]
  spec: 
    # AWS Spec section described below
    regions: 
      - us-east-1
    accounts:
      - id: "account1"
        local_profile: "account1"
    debug: false
```

## AWS Spec

This is the (nested) spec used by the AWS source plugin.

- `regions` ([]string) (default: Empty. Will use all enabled regions)

  Regions to use.

- `accounts` ([][account](#accounts)) (default: current account)

  List of all accounts to fetch information from

- `org` ([org](#org)) (default: not used)

  In AWS organization mode, CloudQuery will source all accounts underneath automatically

- `debug` (bool) (default: false)

  If true, will log AWS debug logs, including retries and other request/response metadata

## accounts

This is used to specify one or more accounts to extract information from. Note that it should be an array of objects, each with the following fields:

- `id` (string) (**required**)

  Will be used as an alias in the source plugin and in the logs

- `local_profile` (string) (default: will use current credentials)

  [Local profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html) to use to authenticate this account with

- `role_arn` (string)

  If specified will use this to assume role

- `role_session_name` (string)

  If specified will use this session name when assume role to `role_arn`

- `external_id` (string)

  If specified will use this when assuming role to `role_arn`

- `default_region` (string) (default: `us-east-1`)

  If specified, this region will be used as the default region for the account.

- `regions` (string)

  Regions to use for this account. Defaults to global `regions` setting.


## org

- `organization_units` ([]string)

  List of Organizational Units that CloudQuery should use to source accounts from

- `admin_account` ([Account](#account))

  Configuration for how to grab credentials from an Admin account

- `member_trusted_principal` ([Account](#account))

  Configuration for how to specify the principle to use in order to assume a role in the member accounts

- `member_role_name` (string) (**required**)

  Role name that CloudQuery should use to assume a role in the member account from the admin account. Note: This is not a full ARN, it is just the name

- `member_role_session_name` (string)

  Override the default Session name.

- `member_external_id` (string)

  Specify an ExternalID for use in the trust policy

- `member_regions` ([]string)

  Limit fetching resources within this specific account to only these regions. This will override any regions specified in the provider block. You can specify all regions by using the `*` character as the only argument in the array
