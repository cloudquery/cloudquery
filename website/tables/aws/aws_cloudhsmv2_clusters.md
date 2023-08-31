# Table: aws_cloudhsmv2_clusters

This table shows data for AWS CloudHSM v2 Clusters.

https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_Cluster.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|backup_policy|`utf8`|
|backup_retention_policy|`json`|
|certificates|`json`|
|cluster_id|`utf8`|
|create_timestamp|`timestamp[us, tz=UTC]`|
|hsm_type|`utf8`|
|hsms|`json`|
|pre_co_password|`utf8`|
|security_group|`utf8`|
|source_backup_id|`utf8`|
|state|`utf8`|
|state_message|`utf8`|
|subnet_mapping|`json`|
|vpc_id|`utf8`|