
# Table: gcp_storage_bucket_policies
A bucket/object IAM policy
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|bucket_id|uuid|Unique ID of gcp_storage_buckets table (FK)|
|etag|text|HTTP 11  Entity tag for the policy|
|kind|text|The kind of item this is For policies, this is always storage#policy This field is ignored on input|
|resource_id|text|The ID of the resource to which this policy belongs Will be of the form projects/_/buckets/bucket for buckets, and projects/_/buckets/bucket/objects/object for objects A specific generation may be specified by appending #generationNumber to the end of the object name, eg projects/_/buckets/my-bucket/objects/datatxt#17 The current generation can be denoted with #0 This field is ignored on input|
|version|bigint|The IAM policy format version|
