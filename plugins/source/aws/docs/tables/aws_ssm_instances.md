
# Table: aws_ssm_instances
Describes a filter for a specific list of instances.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the managed instance.|
|activation_id|text|The activation ID created by Amazon Web Services Systems Manager when the server or virtual machine (VM) was registered.|
|agent_version|text|The version of SSM Agent running on your Linux instance.|
|association_overview_detailed_status|text|Detailed status information about the aggregated associations.|
|association_instance_status_aggregated_count|jsonb|The number of associations for the instance(s).|
|association_status|text|The status of the association.|
|computer_name|text|The fully qualified host name of the managed instance.|
|ip_address|inet|The IP address of the managed instance.|
|iam_role|text|The Identity and Access Management (IAM) role assigned to the on-premises Systems Manager managed instance|
|instance_id|text|The instance ID.|
|is_latest_version|boolean|Indicates whether the latest version of SSM Agent is running on your Linux Managed Instance|
|last_association_execution_date|timestamp without time zone|The date the association was last run.|
|last_ping_date_time|timestamp without time zone|The date and time when the agent last pinged the Systems Manager service.|
|last_successful_association_execution_date|timestamp without time zone|The last date the association was successfully run.|
|name|text|The name assigned to an on-premises server or virtual machine (VM) when it is activated as a Systems Manager managed instance|
|ping_status|text|Connection status of SSM Agent|
|platform_name|text|The name of the operating system platform running on your instance.|
|platform_type|text|The operating system platform type.|
|platform_version|text|The version of the OS platform running on your instance.|
|registration_date|timestamp without time zone|The date the server or VM was registered with Amazon Web Services as a managed instance.|
|resource_type|text|The type of instance|
