# AWS Source Plugin Configuration Reference

## Examples

### Single Account Example

This example connects a single AWS account in one region to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml copy
kind: source
spec:
  # Source spec section
  name: aws
  path: cloudquery/aws
  registry: cloudquery
  version: "VERSION_SOURCE_AWS"
  tables: ["aws_ec2_instances"]
  destinations: ["DESTINATION_NAME"]
  spec:
    # AWS Spec section described below
    regions:
      - us-east-1
    accounts:
      - id: "account1"
        local_profile: "account1"
    aws_debug: false
```

See [tables](/docs/plugins/sources/aws/tables) for a list of all available tables.

### AWS Organization Example

CloudQuery supports discovery of AWS Accounts via AWS Organizations. This means that as Accounts get added or removed from your organization CloudQuery will be able to handle new or removed accounts without any configuration changes.

```yaml copy
kind: source
spec:
  name: aws
  path: cloudquery/aws
  registry: cloudquery
  version: "VERSION_SOURCE_AWS"
  tables: ['aws_s3_buckets']
  destinations: ["DESTINATION_NAME"]
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

- `regions` (`[]string`) (default: `[]`. Will use all enabled regions)

  Regions to use.

- `accounts` ([][account](#account)) (default: current account)

  List of all accounts to fetch information from

- `org` ([org](#org)) (default: not used)

  In AWS organization mode, CloudQuery will source all accounts underneath automatically

- `concurrency` (`int`) (default: `50000`):

  The best effort maximum number of Go routines to use. Lower this number to reduce memory usage.

- `initialization_concurrency` (`int`) (default: `4`)

  During initialization the AWS source plugin fetches information about each account and region. This setting controls how many accounts can be initialized concurrently.
  Only configurations with many accounts (either hardcoded or discovered via Organizations) should require modifying this setting, to either lower it to avoid rate limit errors, or to increase it to speed up the initialization process.

- `scheduler` (`string`) (default: `shuffle`):

  The scheduler to use when determining the priority of resources to sync.

  Currently, the only supported values are `dfs` (depth-first search), `round-robin` and `shuffle`.
  For more information about this, see [performance tuning](/docs/advanced-topics/performance-tuning).

- `aws_debug` (`bool`) (default: `false`)

  If true, will log AWS debug logs, including retries and other request/response metadata

- `max_retries` (`int`) (default: `10`)

  Defines the maximum number of times an API request will be retried

- `max_backoff` (`int`) (default: `30`)

  Defines the duration between retry attempts

- `custom_endpoint_url` (`string`) (default: not used)

  The base URL endpoint the SDK API clients will use to make API calls to. The SDK will suffix URI path and query elements to this endpoint

- `custom_endpoint_hostname_immutable` (`bool`) (default: not used)

  Specifies if the endpoint's hostname can be modified by the SDK's API client.
  When using something like LocalStack make sure to set it equal to `true`.

- `custom_endpoint_partition_id` (`string`) (default: not used)

  The AWS partition the endpoint belongs to

- `custom_endpoint_signing_region` (`string`) (default: not used)

  The region that should be used for signing the request to the endpoint

- `use_paid_apis` (`bool`) (default: `false`)

  When set to `true` plugin will sync data from APIs that incur a fee.

- **enterprise version only** `table_options` (`map`) (default: not used)

- **enterprise version only** `event_based_sync` (`array`) (default: empty)

### account

This is used to specify one or more accounts to extract information from. Note that it should be an array of objects, each with the following fields:

- `id` (string) (**required**)

  Will be used as an alias in the source plugin and in the logs

- `local_profile` (string) (default: will use current credentials)

  [Local profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) to use to authenticate this account with.
  Please note this should be set to the name of the profile.

  For example, with the following credentials file:

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

### org

- `admin_account` ([Account](#account))

  Configuration for how to grab credentials from an Admin account

- `member_trusted_principal` ([Account](#account))

  Configuration for how to specify the principle to use in order to assume a role in the member accounts

- `member_role_name` (string) (**required**)

  Role name that CloudQuery should use to assume a role in the member account from the admin account.

  Note: This is not a full ARN, it is just the name.

- `member_role_session_name` (string)

  Overrides the default session name.

- `member_external_id` (string)

  Specify an external ID for use in the trust policy

- `member_regions` ([]string)

  Limit fetching resources within this specific account to only these regions.
  This will override any regions specified in the provider block.
  You can specify all regions by using the `*` character as the only argument in the array.

- `organization_units` ([]string)

  List of Organizational Units that CloudQuery should use to source accounts from.
 If you specify an OU, CloudQuery will also traverse nested OUs.

- `skip_organization_units` ([]string)

  List of Organizational Units to skip.
 This is useful in conjunction with `organization_units` if there are child OUs that should be ignored.

- `skip_member_accounts` ([]string)

  List of OU member accounts to skip. This is useful if there are accounts under the selected OUs that should be ignored.


<!-- vale off -->
### event_based_sync
<!-- vale on -->

:::callout{type="info"}
Event-based syncing is only supported by the enterprise-version AWS plugin. Read more about [event-based syncs](/docs/plugins/sources/aws/event-based-sync)
:::

- `account` ([Account](#account)) 

  Configuration for the credentials that will be used to grab records from the specified Kinesis Stream. If this is not specified the default credentials will be used

- `kinesis_stream_arn` (string) (**required**)

  ARN for the Kinesis stream that will hold all of the CloudTrail records

- `start_time` (timestamp) 

  Defines the place in the stream where record processing should begin. By default, the time at which the sync began will be used. The value should follow the RFC 3339 format. For example `2023-09-04T19:24:14Z`

- `full_sync` (`bool`) (default: `true`)

  By default, CQ will do a full sync on the specified tables before starting to consume the events in the stream. This parameter enables users to skip the full pull based sync and go straight to the event based sync.


## Advanced Configuration

### Incremental Tables

Some tables, like `aws_cloudtrail_events`, support incremental syncs. When incremental syncing is enabled, CloudQuery will only fetch new data since the last sync. This is useful for tables that have a lot of data and are updated frequently. To enable incremental syncs, add a `backend_options` section to the source config:

```yaml
kind: source
spec:
  # Source spec section
  name: aws
  path: cloudquery/aws
  registry: cloudquery
  version: "VERSION_SOURCE_AWS"
  tables: ["aws_cloudtrail_events"]
  destinations: ["postgresql"]
  backend_options:
    table_name: "cq_state_aws"
    connection: "@@plugins.postgresql.connection"
  spec:
---
kind: destination
spec:
  name: "postgresql"
  path: "cloudquery/postgresql"
  registry: cloudquery
  version: "VERSION_DESTINATION_POSTGRESQL"
  write_mode: "overwrite-delete-stale"
  spec:
    connection_string: "${CONNECTION_STRING}"
```

The `connection` string can reference any destination that supports `overwrite` mode; in the example above it will use the same `postgresql` destination that the `aws_cloudtrail_events` table is written to. The `table_name` is the name of the table that will be used to store state. This table will be created automatically if it does not exist. For more information about managing state for incremental tables, see [Managing Incremental Tables](/docs/advanced-topics/managing-incremental-tables).

### Skip Tables

AWS has tables that may contain many resources, nested information, and AWS-provided data. These tables may cause certain syncs to be slow due to the amount of AWS-provided data and may not be needed. We recommend only specifying syncing from necessary tables. If `*` is necessary for tables, Below is a reference configuration of skip tables, where certain tables are skipped.

```yaml
kind: source
spec:
  # Source spec section
  name: aws
  path: cloudquery/aws
  registry: cloudquery
  version: "VERSION_SOURCE_AWS"
  tables: ["*"]
  skip_tables:
    - aws_cloudtrail_events
    - aws_docdb_cluster_parameter_groups
    - aws_docdb_engine_versions
    - aws_ec2_instance_types
    - aws_ec2_vpc_endpoint_services
    - aws_elasticache_engine_versions
    - aws_elasticache_parameter_groups
    - aws_elasticache_reserved_cache_nodes_offerings
    - aws_elasticache_service_updates
    - aws_iam_group_last_accessed_details
    - aws_iam_policy_last_accessed_details
    - aws_iam_role_last_accessed_details
    - aws_iam_user_last_accessed_details
    - aws_neptune_cluster_parameter_groups
    - aws_neptune_db_parameter_groups
    - aws_rds_cluster_parameter_groups
    - aws_rds_db_parameter_groups
    - aws_rds_engine_versions
    - aws_servicequotas_services
    - aws_stepfunctions_map_run_executions
    - aws_stepfunctions_map_runs
  destinations: ["postgresql"]
  spec:
    # AWS Spec section described below
```
