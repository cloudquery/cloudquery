# Table: aws_ram_resource_share_associations

https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareAssociation.html

The composite primary key for this table is (**associated_entity**, **resource_share_arn**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|associated_entity (PK)|String|
|association_type|String|
|creation_time|Timestamp|
|external|Bool|
|last_updated_time|Timestamp|
|resource_share_arn (PK)|String|
|resource_share_name|String|
|status|String|
|status_message|String|