# Table: aws_redshift_cluster_parameters

This table shows data for Redshift Cluster Parameters.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_Parameter.html

The composite primary key for this table is (**cluster_arn**, **parameter_name**).

## Relations

This table depends on [aws_redshift_cluster_parameter_groups](aws_redshift_cluster_parameter_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_arn (PK)|`utf8`|
|parameter_name (PK)|`utf8`|
|allowed_values|`utf8`|
|apply_type|`utf8`|
|data_type|`utf8`|
|description|`utf8`|
|is_modifiable|`bool`|
|minimum_engine_version|`utf8`|
|parameter_value|`utf8`|
|source|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Connections to Amazon Redshift clusters should be encrypted in transit

```sql
SELECT
  'Connections to Amazon Redshift clusters should be encrypted in transit'
    AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_redshift_clusters AS rsc
WHERE
  EXISTS(
    SELECT
      1
    FROM
      aws_redshift_cluster_parameter_groups AS rscpg
      INNER JOIN aws_redshift_cluster_parameters AS rscp ON
          rscpg.cluster_arn = rscp.cluster_arn
    WHERE
      rsc.arn = rscpg.cluster_arn
      AND (
          rscp.parameter_name = 'require_ssl'
          AND rscp.parameter_value = 'false'
        )
      OR (rscp.parameter_name = 'require_ssl' AND rscp.parameter_value IS NULL)
      OR NOT
          EXISTS(
            (
              SELECT
                1
              FROM
                aws_redshift_cluster_parameters
              WHERE
                cluster_arn = rscpg.cluster_arn
                AND parameter_name = 'require_ssl'
            )
          )
  );
```


