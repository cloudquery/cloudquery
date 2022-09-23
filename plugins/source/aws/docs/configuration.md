# AWS Source Plugin Configuration Reference

## AWS Spec

This is the top level spec used by AWS source plugin.

- `regions` ([]string) (default: Empty. Will use all enabled regions)

  Regions to use.

- `accounts` ([][account](#account)) (default: current account)

  All accounts to fetch information from

- `organization` ([Organization](#organization)) (default: not used)

  In AWS organization mode will source all accounts underneath automatically

- `debug` (bool) (default: false)

  If true will log AWS debug logs including retries and other requests/response metadata

## Account

This is used to specify one or more accounts to extract information from.

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
