# Alibaba Cloud Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `alicloud`)}/>

The Alibaba Cloud source plugin for CloudQuery extracts configuration from the [Alibaba Cloud (阿里云) API](https://www.alibabacloud.com/product/openapiexplorer) and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Configuration

The following configuration syncs from Alibaba Cloud to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](https://www.cloudquery.io/docs/reference/source-spec). The config for the `postgresql` destination is not shown here. See our [Quickstart](/docs/quickstart) if you need help setting up the destination.

```yaml
kind: source
spec:
  name: "alicloud"
  path: "cloudquery/alicloud"
  version: "VERSION_SOURCE_ALICLOUD"
  tables: ["*"]
  destinations: 
    - "postgresql"
  spec:
    accounts: 
      - name: my_account
        regions:
        - cn-hangzhou
        - cn-beijing
        - eu-west-1
        - us-west-1
        # ...
        access_key: ${ALICLOUD_ACCESS_KEY}
        secret_key: ${ALICLOUD_SECRET_KEY}
```

- `accounts` (array[object], required):

  A list of accounts to sync. Every account must have a unique name, and must specify at least one region. The `access_key` and `secret_key` are required and can be specified as environment variables, as shown in the example above.
  - `name` (string, required): A unique name for the account.
  - `regions` (array[string], required): A list of regions to sync. For example, `["cn-hangzhou", "cn-beijing"]`.
  - `access_key` (string, required): A valid access key for the account
  - `secret_key` (string, required): A valid secret key for the account, corresponding to the access key

- `bill_history_months` (int, optional):

  The number of months of billing history to fetch for the `alicloud_bss_bill` and `alicloud_bss_bill_overview` tables. Defaults to 12.


See the [Alibaba documentation](https://www.alibabacloud.com/help/en/basics-for-beginners/latest/obtain-an-accesskey-pair) for how to obtain an AccessKey pair.

## Example Queries

### Find all ECS instances in a region

```sql
select 
  instance_id, 
  os_name, 
  region_id, 
  start_time, 
  tags 
from 
  alicloud_ecs_instances 
where 
  region_id = 'eu-west-1';
```

```text
+------------------------+--------------------------------------+-----------+-------------------+---------------+
| instance_id            | os_name                              | region_id | start_time        | tags          |
|------------------------+--------------------------------------+-----------+-------------------+---------------|
| i-xxxxxxxxxxxxxxxxxxxx | Alibaba Cloud Linux  3.2104 LTS 64位 | eu-west-1 | 2023-01-17T14:40Z | {"Tag": null} |
+------------------------+--------------------------------------+-----------+-------------------+---------------+
```

### Query past bills

```sql
select 
  product_name, 
  item, 
  pip_code, 
  currency, 
  adjust_amount 
from 
  alicloud_bss_bill_overview;
```

```text
+------------------------+----------------+----------+----------+---------------+
| product_name           | item           | pip_code | currency | adjust_amount |
|------------------------+----------------+----------+----------+---------------|
| Object Storage Service | PayAsYouGoBill | oss      | USD      | 0.0           |
+------------------------+----------------+----------+----------+---------------+
```

### Query bucket stats

```sql
select 
  account_id, 
  bucket_name, 
  object_count, 
  storage 
from 
  alicloud_oss_bucket_stats;
```

```text
+------------+-------------+--------------+---------+
| account_id | bucket_name | object_count | storage |
|------------+-------------+--------------+---------|
| test       | cq-test     | 2            | 29665   |
+------------+-------------+--------------+---------+
```