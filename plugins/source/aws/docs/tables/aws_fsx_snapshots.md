
# Table: aws_fsx_snapshots
A snapshot of an Amazon FSx for OpenZFS volume.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|creation_time|timestamp without time zone|The time that the resource was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time.|
|lifecycle|text|The lifecycle status of the snapshot.  * PENDING - Amazon FSx hasn't started creating the snapshot.  * CREATING - Amazon FSx is creating the snapshot.  * DELETING - Amazon FSx is deleting the snapshot.  * AVAILABLE - The snapshot is fully available.|
|lifecycle_transition_reason_message|text|A detailed error message.|
|name|text|The name of the snapshot.|
|arn|text|The Amazon Resource Name (ARN) for a given resource|
|snapshot_id|text|The ID of the snapshot.|
|tags|jsonb|A list of Tag values, with a maximum of 50 elements.|
|volume_id|text|The ID of the volume that the snapshot is of.|
