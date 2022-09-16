
# Table: aws_lightsail_static_ips
Describes a static IP
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the static IP (eg, arn:aws:lightsail:us-east-2:123456789101:StaticIp/9cbb4a9e-f8e3-4dfe-b57e-12345EXAMPLE)|
|attached_to|text|The instance where the static IP is attached (eg, Amazon_Linux-1GB-Ohio-1)|
|created_at|timestamp without time zone|The timestamp when the static IP was created (eg, 1479735304222)|
|ip_address|text|The static IP address|
|is_attached|boolean|A Boolean value indicating whether the static IP is attached|
|availability_zone|text|The Availability Zone|
|name|text|The name of the static IP (eg, StaticIP-Ohio-EXAMPLE)|
|resource_type|text|The resource type (usually StaticIp)|
|support_code|text|The support code|
