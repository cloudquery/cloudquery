
# Table: aws_fsx_volumes
Describes an Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS volume
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|creation_time|timestamp without time zone|The time that the resource was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time|
|file_system_id|text|The globally unique ID of the file system, assigned by Amazon FSx|
|lifecycle|text|The lifecycle status of the volume  * AVAILABLE - The volume is fully available for use  * CREATED - The volume has been created  * CREATING - Amazon FSx is creating the new volume  * DELETING - Amazon FSx is deleting an existing volume  * FAILED - Amazon FSx was unable to create the volume  * MISCONFIGURED - The volume is in a failed but recoverable state  * PENDING - Amazon FSx hasn't started creating the volume|
|lifecycle_transition_reason_message|text|A detailed error message|
|name|text|The name of the volume|
|arn|text|The Amazon Resource Name (ARN) for a given resource|
|tags|jsonb|A list of Tag values, with a maximum of 50 elements|
|id|text|The system-generated, unique ID of the volume|
|volume_type|text|The type of the volume|
