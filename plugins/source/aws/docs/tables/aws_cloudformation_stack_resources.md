# Table: aws_cloudformation_stack_resources

https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackResourceSummary.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_cloudformation_stacks](aws_cloudformation_stacks.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|last_updated_timestamp|Timestamp|
|logical_resource_id|String|
|resource_status|String|
|resource_type|String|
|drift_information|JSON|
|module_info|JSON|
|physical_resource_id|String|
|resource_status_reason|String|