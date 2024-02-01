# Table: aws_elasticsearch_domains

This table shows data for Elasticsearch Domains.

https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_DomainStatus.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|authorized_principals|`json`|
|tags|`json`|
|arn|`utf8`|
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