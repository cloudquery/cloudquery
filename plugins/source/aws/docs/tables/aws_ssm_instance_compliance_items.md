
# Table: aws_ssm_instance_compliance_items
Information about the compliance as defined by the resource type
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_ssm_instances table (FK)|
|compliance_type|text|The compliance type|
|details|jsonb|A "Key": "Value" tag combination for the compliance item.|
|execution_summary_execution_time|timestamp without time zone|The time the execution ran as a datetime object that is saved in the following format: yyyy-MM-dd'T'HH:mm:ss'Z'.|
|execution_summary_execution_id|text|An ID created by the system when PutComplianceItems was called|
|execution_summary_execution_type|text|The type of execution|
|id|text|An ID for the compliance item|
|resource_id|text|An ID for the resource|
|resource_type|text|The type of resource|
|severity|text|The severity of the compliance status|
|status|text|The status of the compliance item|
|title|text|A title for the compliance item|
