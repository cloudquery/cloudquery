
# Table: aws_ec2_images
Describes an image.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|id|text|The ID of the AMI.|
|architecture|text|The architecture of the image.|
|creation_date|text|The date and time the image was created.|
|description|text|The description of the AMI that was provided during image creation.|
|ena_support|boolean|Specifies whether enhanced networking with ENA is enabled.|
|hypervisor|text|The hypervisor type of the image.|
|image_location|text|The location of the AMI.|
|image_owner_alias|text|The AWS account alias (for example, amazon, self) or the AWS account ID of the AMI owner.|
|image_type|text|The type of image.|
|kernel_id|text|The kernel associated with the image, if any.|
|name|text|The name of the AMI that was provided during image creation.|
|owner_id|text|The AWS account ID of the image owner.|
|platform|text|This value is set to windows for Windows AMIs; otherwise, it is blank.|
|platform_details|text|The platform details associated with the billing code of the AMI.|
|product_codes|jsonb|Any product codes associated with the AMI.|
|public|boolean|Indicates whether the image has public launch permissions.|
|ramdisk_id|text|The RAM disk associated with the image, if any.|
|root_device_name|text|The device name of the root device volume (for example, /dev/sda1).|
|root_device_type|text|The type of root device used by the AMI.|
|sriov_net_support|text|Specifies whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.|
|state|text|The current state of the AMI.|
|state_reason_code|text|The reason code for the state change.|
|state_reason_message|text|The message for the state change.|
|tags|jsonb|Any tags assigned to the image.|
|usage_operation|text|The operation of the Amazon EC2 instance and the billing code that is associated with the AMI.|
|virtualization_type|text|The type of virtualization of the AMI.|
|last_launched_time|timestamp without time zone|The timestamp of the last time the AMI was used for an EC2 instance launch.|
