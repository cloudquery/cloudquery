# AWS Source Plugin Configuration Reference

## Simple Example

This example connects a single AWS account in one region to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml copy
kind: source
spec:
  # Source spec section
  name: aws
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  tables: ["*"]
  destinations: ["postgresql"]
  spec: 
    # AWS Spec section described below
    regions: 
      - us-east-1
    accounts:
      - id: "account1"
        local_profile: "account1"
    aws_debug: false
```

## Skipping tables with configuration parameters

Some tables document the parameters and options available to your AWS accounts and don't correspond to real resources. If you don't need these tables, the time it takes to sync can be reduced by skipping these tables: 

```yaml copy
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  tables: ["*"]

  # Comment out any of the following tables if you want to sync them
  # unless otherwise indicated they are configuration parameters rather than configured resources
  skip_tables:
    - aws_ec2_vpc_endpoint_services # this resource includes services that are available from AWS as well as other AWS Accounts
    - aws_docdb_cluster_parameter_groups
    - aws_docdb_engine_versions
    - aws_ec2_instance_types
    - aws_elasticache_engine_versions
    - aws_elasticache_parameter_groups
    - aws_elasticache_reserved_cache_nodes_offerings
    - aws_elasticache_service_updates
    - aws_neptune_cluster_parameter_groups
    - aws_neptune_db_parameter_groups
    - aws_rds_cluster_parameter_groups
    - aws_rds_db_parameter_groups
    - aws_rds_engine_versions
    - aws_servicequotas_services
  destinations: ["<destination>"]
``` 

## AWS Organization Example

CloudQuery supports discovery of AWS Accounts via AWS Organizations. This means that as Accounts get added or removed from your organization CloudQuery will be able to handle new or removed accounts without any configuration changes.

```yaml copy
kind: source
spec:
  name: aws
  registry: github
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  tables: ['*']
  destinations: ["postgresql"]
  spec:
    aws_debug: false
    org:
      admin_account:
        local_profile: "<NAMED_PROFILE>"
      member_role_name: OrganizationAccountAccessRole
    regions:
      - '*'
  ```

For full details, see the [Multi Account Configuration Tutorial](/docs/plugins/sources/aws/multi-account).

## AWS Spec

This is the (nested) spec used by the AWS source plugin.

- `regions` ([]string) (default: Empty. Will use all enabled regions)

  Regions to use.

- `accounts` ([][account](#accounts)) (default: current account)

  List of all accounts to fetch information from

- `org` ([org](#org)) (default: not used)

  In AWS organization mode, CloudQuery will source all accounts underneath automatically

- `aws_debug` (bool) (default: false)

  If true, will log AWS debug logs, including retries and other request/response metadata

- `max_retries` (int) (default: 10)

  Defines the maximum number of times an API request will be retried 

- `max_backoff` (int) (default: 30)
  
  Defines the duration between retry attempts

- `custom_endpoint_url` (string) (default: not used)

  The base URL endpoint the SDK API clients will use to make API calls to. The SDK will suffix URI path and query elements to this endpoint

- `custom_endpoint_hostname_immutable` (bool) (default: not used)

  Specifies if the endpoint's hostname can be modified by the SDK's API client. When using something like LocalStack make sure to set it equal to `True`

- `custom_endpoint_partition_id` (string) (default: not used)

  The AWS partition the endpoint belongs to

- `custom_endpoint_signing_region` (string) (default: not used)

  The region that should be used for signing the request to the endpoint


## accounts

This is used to specify one or more accounts to extract information from. Note that it should be an array of objects, each with the following fields:

- `id` (string) (**required**)

  Will be used as an alias in the source plugin and in the logs

- `local_profile` (string) (default: will use current credentials)

  [Local profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html) to use to authenticate this account with.
  Please note this should be set to the name of the profile. For example, with the following credentials file:

  ```ini copy
  [default]
  aws_access_key_id=xxxx
  aws_secret_access_key=xxxx

  [user1]
  aws_access_key_id=xxxx
  aws_secret_access_key=xxxx
  ```

  `local_profile` should be set to either `default` or `user1`.

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

- `organization_units` ([]string)

  List of Organizational Units that CloudQuery should use to source accounts from. If you specify an OU, CloudQuery will not traverse nested OUs

- `skip_organization_units` ([]string)

  List of Organizational Units to skip. This is useful in conjunction with `organization_units` if there are child OUs that should be ignored.

- `skip_member_accounts` ([]string)

  List of OU member accounts to skip. This is useful in conjunction with `organization_units` if there are accounts under the selected OUs that should be ignored.
