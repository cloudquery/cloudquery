# Table: aws_lightsail_container_service_deployments

This table shows data for Lightsail Container Service Deployments.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerServiceDeployment.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_lightsail_container_services](aws_lightsail_container_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|container_service_arn|`utf8`|
|containers|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|public_endpoint|`json`|
|state|`utf8`|
|version|`int64`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Unused Lightsail container services

```sql
WITH
  deployment
    AS (
      SELECT
        DISTINCT container_service_arn
      FROM
        aws_lightsail_container_service_deployments
    )
SELECT
  'Unused Lightsail container services' AS title,
  cs.account_id,
  cs.arn AS resource_id,
  'fail' AS status
FROM
  aws_lightsail_container_services AS cs
  LEFT JOIN deployment ON deployment.container_service_arn = cs.arn
WHERE
  deployment.container_service_arn IS NULL;
```


