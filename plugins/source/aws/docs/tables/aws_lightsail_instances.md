
# Table: aws_lightsail_instances
Describes an instance (a virtual private server)
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|access_details|jsonb||
|arn|text|The Amazon Resource Name (ARN) of the instance (eg, arn:aws:lightsail:us-east-2:123456789101:Instance/244ad76f-8aad-4741-809f-12345EXAMPLE)|
|blueprint_id|text|The blueprint ID (eg, os_amlinux_2016_03)|
|blueprint_name|text|The friendly name of the blueprint (eg, Amazon Linux)|
|bundle_id|text|The bundle for the instance (eg, micro_1_0)|
|created_at|timestamp without time zone|The timestamp when the instance was created (eg, 147973490917) in Unix time format|
|hardware_cpu_count|bigint|The number of vCPUs the instance has|
|hardware_ram_size_in_gb|float|The amount of RAM in GB on the instance (eg, 10)|
|ip_address_type|text|The IP address type of the instance|
|ipv6_addresses|text[]|The IPv6 addresses of the instance|
|is_static_ip|boolean|A Boolean value indicating whether this instance has a static IP assigned to it|
|location_availability_zone|text|The Availability Zone|
|location_region_name|text|The AWS Region name|
|name|text|The name the user gave the instance (eg, Amazon_Linux-1GB-Ohio-1)|
|networking_monthly_transfer_gb_per_month_allocated|bigint|The amount allocated per month (in GB)|
|private_ip_address|text|The private IP address of the instance|
|public_ip_address|text|The public IP address of the instance|
|resource_type|text|The type of resource (usually Instance)|
|ssh_key_name|text|The name of the SSH key being used to connect to the instance (eg, LightsailDefaultKeyPair)|
|state_code|bigint|The status code for the instance|
|state_name|text|The state of the instance (eg, running or pending)|
|support_code|text|The support code|
|tags|jsonb|The tag keys and optional values for the resource|
|username|text|The user name for connecting to the instance (eg, ec2-user)|
