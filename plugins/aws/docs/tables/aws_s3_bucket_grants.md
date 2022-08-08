
# Table: aws_s3_bucket_grants
Container for grant information.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|bucket_cq_id|uuid|Unique CloudQuery ID of aws_s3_buckets table (FK)|
|type|text|Type of grantee|
|display_name|text|Screen name of the grantee.|
|email_address|text|Email address of the grantee|
|grantee_id|text|The canonical user ID of the grantee.|
|uri|text|URI of the grantee group.|
|permission|text|Specifies the permission given to the grantee.|
