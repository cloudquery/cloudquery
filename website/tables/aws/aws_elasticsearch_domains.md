# Table: aws_elasticsearch_domains

This table shows data for Elasticsearch Domains.

https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_DomainStatus.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|authorized_principals|`json`|
|tags|`json`|
|arn (PK)|`utf8`|
|domain_id|`utf8`|
|domain_name|`utf8`|
|elasticsearch_cluster_config|`json`|
|access_policies|`utf8`|
|advanced_options|`json`|
|advanced_security_options|`json`|
|auto_tune_options|`json`|
|change_progress_details|`json`|
|cognito_options|`json`|
|created|`bool`|
|deleted|`bool`|
|domain_endpoint_options|`json`|
|ebs_options|`json`|
|elasticsearch_version|`utf8`|
|encryption_at_rest_options|`json`|
|endpoint|`utf8`|
|endpoints|`json`|
|log_publishing_options|`json`|
|node_to_node_encryption_options|`json`|
|processing|`bool`|
|service_software_options|`json`|
|snapshot_options|`json`|
|upgrade_processing|`bool`|
|vpc_options|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Connections to Elasticsearch domains should be encrypted using TLS 1.2

```sql
SELECT
  'Connections to Elasticsearch domains should be encrypted using TLS 1.2'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN domain_endpoint_options->>'TLSSecurityPolicy'
  IS DISTINCT FROM 'Policy-Min-TLS-1-2-2019-07'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elasticsearch_domains;
```

### Elasticsearch domain error logging to CloudWatch Logs should be enabled

```sql
SELECT
  'Elasticsearch domain error logging to CloudWatch Logs should be enabled'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN log_publishing_options->'ES_APPLICATION_LOGS'->'Enabled'
  IS DISTINCT FROM 'true'
  OR (
      log_publishing_options->'ES_APPLICATION_LOGS'->'CloudWatchLogsLogGroupArn'
    ) IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elasticsearch_domains;
```

### Elasticsearch domains should be configured with at least three dedicated master nodes

```sql
SELECT
  'Elasticsearch domains should be configured with at least three dedicated master nodes'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (elasticsearch_cluster_config->>'DedicatedMasterEnabled')::BOOL
  IS NOT true
  OR (elasticsearch_cluster_config->>'DedicatedMasterCount')::INT8 IS NULL
  OR (elasticsearch_cluster_config->>'DedicatedMasterCount')::INT8 < 3
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elasticsearch_domains;
```

### Elasticsearch domains should be in a VPC

```sql
SELECT
  'Elasticsearch domains should be in a VPC' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (vpc_options->>'VPCId') IS NULL THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elasticsearch_domains;
```

### Elasticsearch domains should encrypt data sent between nodes

```sql
SELECT
  'Elasticsearch domains should encrypt data sent between nodes' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (node_to_node_encryption_options->>'Enabled')::BOOL IS NOT true
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elasticsearch_domains;
```

### Elasticsearch domains should have at least three data nodes

```sql
SELECT
  'Elasticsearch domains should have at least three data nodes' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN NOT (elasticsearch_cluster_config->>'ZoneAwarenessEnabled')::BOOL
  OR (elasticsearch_cluster_config->>'InstanceCount')::INT8 IS NULL
  OR (elasticsearch_cluster_config->>'InstanceCount')::INT8 < 3
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elasticsearch_domains;
```

### Elasticsearch domains should have audit logging enabled

```sql
SELECT
  'Elasticsearch domains should have audit logging enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN log_publishing_options->'AUDIT_LOGS'->'Enabled' IS DISTINCT FROM 'true'
  OR (log_publishing_options->'AUDIT_LOGS'->'CloudWatchLogsLogGroupArn') IS NULL
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elasticsearch_domains;
```

### Elasticsearch domains should have encryption at rest enabled

```sql
SELECT
  'Elasticsearch domains should have encryption at rest enabled' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN (encryption_at_rest_options->>'Enabled')::BOOL IS NOT true THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_elasticsearch_domains;
```


