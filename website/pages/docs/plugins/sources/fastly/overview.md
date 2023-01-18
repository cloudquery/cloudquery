# Fastly Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `fastly`)}/>

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

- `services` ([]string, optional):

  A list of Fastly service IDs to sync. If not specified, all services will be used.

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

### Retrieve stats for a service for a specific time period

```sql copy
SELECT   to_char(Date_trunc('month', start_time), 'Month')             AS month,
         sum(requests)                                                 AS requests,
         pg_size_pretty(Sum(resp_body_bytes) + sum(resp_header_bytes)) AS resp_bytes,
         sum(status_ 2xx)                                              AS status_2xx,
         sum(status_4xx)                                               AS status_4xx,
         sum(status_5xx)                                               AS status_5xx
FROM     fastly_stats_services
WHERE    service_id = '1234567890abcdefghijkl'
AND      start_time >= date '2022-01-01'
AND      s tart_time < date '2023-01-01'
GROUP BY date_trunc('month', start_time)
ORDER BY date_trunc('month', start_time) ASC
```

```text
+-----------+----------+------------+------------+------------+------------+
| month     | requests | resp_bytes | status_2xx | status_4xx | status_5xx |
|-----------+----------+------------+------------+------------+------------|
| January   | 24274    | 225 MB     | 17526      | 1937       | 43         |
| February  | 26584    | 251 MB     | 17817      | 4232       | 14         |
| March     | 24508    | 240 MB     | 18416      | 1788       | 18         |
| April     | 25098    | 243 MB     | 17892      | 3066       | 142        |
| May       | 25865    | 254 MB     | 18647      | 2849       | 18         |
| June      | 18001    | 181 MB     | 12487      | 2711       | 5          |
| July      | 22005    | 206 MB     | 14759      | 3414       | 30         |
| August    | 19737    | 186 MB     | 12824      | 3344       | 14         |
| September | 24001    | 235 MB     | 15944      | 4483       | 5          |
| October   | 23244    | 240 MB     | 16180      | 3099       | 8          |
| November  | 22119    | 201 MB     | 15237      | 3832       | 2          |
| December  | 18767    | 180 MB     | 13414      | 2423       | 18         |
+-----------+----------+------------+------------+------------+------------+
```

### Select users that don't have 2FA enabled

```sql copy
SELECT name,
       login,
       two_factor_auth_enabled
FROM   fastly_account_users
WHERE  two_factor_auth_enabled IS FALSE
```

```text
+---------------+---------------------------+-------------------------+
| name          | login                     | two_factor_auth_enabled |
|---------------+---------------------------+-------------------------|
| Rudolph       | rudolph@gmail.com         | False                   |
| Santa         | santa.northpole@gmail.com | False                   |
+---------------+---------------------------+-------------------------+
```

### List all API tokens and their expiry dates

```sql copy
SELECT NAME,
       scope,
       created_at,
       last_used_at,
       expires_at
FROM   fastly_auth_tokens 
```

```text
+-----------------------------------+-------------+---------------------+---------------------+---------------------+
| name                              | scope       | created_at          | last_used_at        | expires_at          |
|-----------------------------------+-------------+---------------------+---------------------+---------------------|
| manage.fastly.com browser session | global      | 2022-12-26 12:24:01 | <null>              | <null>              |
| Robot Santa Claus                 | global:read | 2022-11-25 00:00:00 | <null>              | 2024-12-25 00:00:00 |
+-----------------------------------+-------------+---------------------+---------------------+---------------------+
```