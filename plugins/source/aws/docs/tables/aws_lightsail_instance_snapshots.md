
# Table: aws_lightsail_instance_snapshots
Describes an instance snapshot
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the snapshot (eg, arn:aws:lightsail:us-east-2:123456789101:InstanceSnapshot/d23b5706-3322-4d83-81e5-12345EXAMPLE)|
|created_at|timestamp without time zone|The timestamp when the snapshot was created (eg, 1479907467024)|
|from_blueprint_id|text|The blueprint ID from which you created the snapshot (eg, os_debian_8_3)|
|from_bundle_id|text|The bundle ID from which you created the snapshot (eg, micro_1_0)|
|from_instance_arn|text|The Amazon Resource Name (ARN) of the instance from which the snapshot was created (eg, arn:aws:lightsail:us-east-2:123456789101:Instance/64b8404c-ccb1-430b-8daf-12345EXAMPLE)|
|from_instance_name|text|The instance from which the snapshot was created|
|is_from_auto_snapshot|boolean|A Boolean value indicating whether the snapshot was created from an automatic snapshot|
|availability_zone|text|The Availability Zone|
|name|text|The name of the snapshot|
|progress|text|The progress of the snapshot|
|resource_type|text|The type of resource (usually InstanceSnapshot)|
|size_in_gb|bigint|The size in GB of the SSD|
|state|text|The state the snapshot is in|
|support_code|text|The support code|
|tags|jsonb|The tag keys and optional values for the resource|
