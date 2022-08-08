
# Table: aws_s3_buckets
An Amazon S3 bucket is a public cloud storage resource available in Amazon Web Services' (AWS) Simple Storage Service (S3)
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|logging_target_prefix|text||
|logging_target_bucket|text||
|versioning_status|text||
|versioning_mfa_delete|text||
|policy|jsonb||
|tags|jsonb||
|creation_date|timestamp without time zone|Date the bucket was created|
|name|text|The name of the bucket.|
|block_public_acls|boolean|Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket|
|block_public_policy|boolean|Specifies whether Amazon S3 should block public bucket policies for this bucket. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access|
|ignore_public_acls|boolean|Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket|
|restrict_public_buckets|boolean|Specifies whether Amazon S3 should restrict public bucket policies for this bucket|
|replication_role|text|The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that Amazon S3 assumes when replicating objects|
|arn|text|The Amazon Resource Name (ARN) for the resource.|
|ownership_controls|text[]|The OwnershipControls (BucketOwnerEnforced, BucketOwnerPreferred, or ObjectWriter) currently in effect for this Amazon S3 bucket.|
