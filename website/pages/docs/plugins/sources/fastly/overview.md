# Fastly Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", "fastly")}/>

The CloudQuery Fastly plugin reads information from your Fastly account(s) and loads it into any supported CloudQuery destination (e.g. PostgreSQL, Snowflake, BigQuery, etc).

## Configuration

This example syncs from Fastly to a Postgres destination, using the token provided in an environment variable called `FASTLY_API_KEY`. The (top level) source spec section is described in the [Source Spec Reference](/docs/reference/source-spec).

```yaml
kind: source
# Common source-plugin configuration
spec:
  name: fastly
  path: cloudquery/fastly
  version: "VERSION_SOURCE_FASTLY"
  tables: ["*"]
  destinations: ["postgresql"]

  # Fastly specific configuration
  spec:
    fastly_api_key: "${FASTLY_API_KEY}"
```

For more information on downloading, installing and running the CloudQuery CLI, see the [Quickstart guide](/docs/quickstart).

## Fastly Spec

This is the (nested) spec used by the Fastly source plugin.

- `fastly_api_key` (string, required):
   
  An API token to access Fastly resources. This can be obtained by [creating a Fastly API token](https://docs.fastly.com/en/guides/using-api-tokens). It should be a *User token* with the *Global:read* scope. Automation tokens do not allow the listing of service versions.

## Example Queries

### List all services and their active versions

```sql copy
select name, id, type, active_version from fastly_services;
```

```text
+----------------------+------------------------+------+----------------+
| name                 | id                     | type | active_version |
|----------------------+------------------------+------+----------------|
| My Test Service      | 1234567890abcdefghijkl | vcl  | 6              |
| Another Service      | 0987654321mnopqrstuvwx | vcl  | 7              |
+----------------------+------------------------+------+----------------+
```

### List domains attached to active versions of all services

```sql copy
SELECT s.id   AS service_id,
       s.name AS service_name,
       s.active_version,
       d.name AS domain_name
FROM fastly_service_domains d
   JOIN fastly_services s
     ON d.service_id = s.id
     AND d.service_version = s.active_version
ORDER BY service_name;
```

```text
+------------------------+----------------------+----------------+--------------------------------------+
| service_id             | service_name         | active_version | domain_name                          |
|------------------------+----------------------+----------------+--------------------------------------|
| 1234567890abcdefghijkl | My Service           | 7              | www.my-service-domain.com            |
| 1234567890abcdefghijkl | My Service           | 7              | my-service-domain.com                |
| 0987654321mnopqrstuvwx | Test Service         | 6              | my-test-service-domain.com           |
+------------------------+----------------------+----------------+--------------------------------------+
```

### Discover how health check definitions changed between service versions

```sql copy
SELECT host,
       path,
       method,
       threshold,
       service_version,
       check_interval
FROM   fastly_service_health_checks
WHERE  service_id = '1234567890abcdefghijkl'
ORDER  BY service_version DESC; 
```

```text
+-----------------------+--------+--------+-----------+-----------------+----------------+
| host                  | path   | method | threshold | service_version | check_interval |
|-----------------------+--------+--------+-----------+-----------------+----------------|
| my-service-domain.com | /blog  | HEAD   | 1         | 4               | 10000          |
| my-service-domain.com | /blog  | HEAD   | 2         | 3               | 20000          |
| my-service-domain.com | /blog  | GET    | 3         | 2               | 30000          |
| my-service-domain.com | /blog  | GET    | 10        | 1               | 40000          |
+-----------------------+--------+--------+-----------+-----------------+----------------+
```