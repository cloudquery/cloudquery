
# Table: aws_cloudformation_stack_resources
Contains high-level information about the specified stack resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|stack_cq_id|uuid|Unique CloudQuery ID of aws_cloudformation_stacks table (FK)|
|last_updated_timestamp|timestamp without time zone|Time the status was updated.  This member is required.|
|logical_resource_id|text|The logical name of the resource specified in the template.  This member is required.|
|resource_status|text|Current status of the resource.  This member is required.|
|resource_type|text|Type of resource|
|stack_resource_drift_status|text|Status of the resource's actual configuration compared to its expected configuration.  * DELETED: The resource differs from its expected configuration in that it has been deleted.  * MODIFIED: The resource differs from its expected configuration.  * NOT_CHECKED: CloudFormation hasn't checked if the resource differs from its expected configuration|
|drift_last_check_timestamp|timestamp without time zone|When CloudFormation last checked if the resource had drifted from its expected configuration.|
|module_info_logical_id_hierarchy|text|A concatenated list of the logical IDs of the module or modules containing the resource|
|module_info_type_hierarchy|text|A concatenated list of the module type or types containing the resource|
|physical_resource_id|text|The name or unique identifier that corresponds to a physical instance ID of the resource.|
|resource_status_reason|text|Success/failure message associated with the resource.|
