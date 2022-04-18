
# Table: aws_ec2_hosts
Describes the properties of the Dedicated Host.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the dedicated host.|
|allocation_time|timestamp without time zone|The time that the Dedicated Host was allocated.|
|allows_multiple_instance_types|text|Indicates whether the Dedicated Host supports multiple instance types of the same instance family|
|auto_placement|text|Whether auto-placement is on or off.|
|availability_zone|text|The Availability Zone of the Dedicated Host.|
|availability_zone_id|text|The ID of the Availability Zone in which the Dedicated Host is allocated.|
|available_vcpus|integer|The number of vCPUs available for launching instances onto the Dedicated Host.|
|client_token|text|Unique, case-sensitive identifier that you provide to ensure the idempotency of the request|
|id|text|The ID of the Dedicated Host.|
|cores|integer|The number of cores on the Dedicated Host.|
|instance_family|text|The instance family supported by the Dedicated Host|
|instance_type|text|The instance type supported by the Dedicated Host|
|sockets|integer|The number of sockets on the Dedicated Host.|
|total_vcpus|integer|The total number of vCPUs on the Dedicated Host.|
|host_recovery|text|Indicates whether host recovery is enabled or disabled for the Dedicated Host.|
|reservation_id|text|The reservation ID of the Dedicated Host|
|member_of_service_linked_resource_group|boolean|Indicates whether the Dedicated Host is in a host resource group|
|owner_id|text|The ID of the Amazon Web Services account that owns the Dedicated Host.|
|release_time|timestamp without time zone|The time that the Dedicated Host was released.|
|state|text|The Dedicated Host's state.|
|tags|jsonb|Any tags assigned to the Dedicated Host.|
