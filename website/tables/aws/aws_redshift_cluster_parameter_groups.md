# Table: aws_redshift_cluster_parameter_groups

This table shows data for Redshift Cluster Parameter Groups.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_ClusterParameterGroupStatus.html

The composite primary key for this table is (**cluster_arn**, **parameter_group_name**).

## Relations

This table depends on [aws_redshift_clusters](aws_redshift_clusters).

The following tables depend on aws_redshift_cluster_parameter_groups:
  - [aws_redshift_cluster_parameters](aws_redshift_cluster_parameters)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_arn (PK)|`utf8`|
|parameter_group_name (PK)|`utf8`|
|cluster_parameter_status_list|`json`|
|parameter_apply_status|`utf8`|

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


