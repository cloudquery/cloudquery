
# Table: aws_s3_bucket_lifecycles
A lifecycle rule for individual objects in an Amazon S3 bucket.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|bucket_id|uuid|Unique ID of aws_s3_buckets table (FK)|
|status|text|If 'Enabled', the rule is currently being applied|
|abort_incomplete_multipart_upload_days_after_initiation|integer|Specifies the number of days after which Amazon S3 aborts an incomplete multipart upload.|
|expiration_date|timestamp without time zone|Indicates at what date the object is to be moved or deleted|
|expiration_days|integer|Indicates the lifetime, in days, of the objects that are subject to the rule. The value must be a non-zero positive integer.|
|expiration_expired_object_delete_marker|boolean|Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions|
|filter|jsonb|The Filter is used to identify objects that a Lifecycle Rule applies to|
|resource_id|text|Unique identifier for the rule|
|noncurrent_version_expiration_days|integer|Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action|
|noncurrent_version_transitions|jsonb|Specifies the transition rule for the lifecycle rule that describes when noncurrent objects transition to a specific storage class|
|prefix|text|Prefix identifying one or more objects to which the rule applies|
|transitions|jsonb|Specifies when an Amazon S3 object transitions to a specified storage class.|
