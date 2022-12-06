# Table: aws_elasticsearch_domains



The composite primary key for this table is (**account_id**, **region**, **id**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|tags|JSON|
|id (PK)|String|
|arn|String|
|domain_name|String|
|elasticsearch_cluster_config|JSON|
|access_policies|String|
|advanced_options|JSON|
|advanced_security_options|JSON|
|auto_tune_options|JSON|
|change_progress_details|JSON|
|cognito_options|JSON|
|created|Bool|
|deleted|Bool|
|domain_endpoint_options|JSON|
|ebs_options|JSON|
|elasticsearch_version|String|
|encryption_at_rest_options|JSON|
|endpoint|String|
|endpoints|JSON|
|log_publishing_options|JSON|
|node_to_node_encryption_options|JSON|
|processing|Bool|
|service_software_options|JSON|
|snapshot_options|JSON|
|upgrade_processing|Bool|
|vpc_options|JSON|