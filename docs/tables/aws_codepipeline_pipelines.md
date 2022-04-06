
# Table: aws_codepipeline_pipelines
Represents the output of a GetPipeline action.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|The tags associated with the pipeline.|
|created|timestamp without time zone|The date and time the pipeline was created, in timestamp format.|
|arn|text|The Amazon Resource Name (ARN) of the pipeline.|
|updated|timestamp without time zone|The date and time the pipeline was last updated, in timestamp format.|
|name|text|The name of the pipeline.  This member is required.|
|role_arn|text|The Amazon Resource Name (ARN) for AWS CodePipeline to use to either perform actions with no actionRoleArn, or to use to assume roles for actions with an actionRoleArn.|
|artifact_store_location|text|The S3 bucket used for storing the artifacts for a pipeline|
|artifact_store_type|text|The type of the artifact store, such as S3.  This member is required.|
|artifact_store_encryption_key_id|text|The ID used to identify the key|
|artifact_store_encryption_key_type|text|The type of encryption key, such as an AWS Key Management Service (AWS KMS) key. When creating or updating a pipeline, the value must be set to 'KMS'.|
|artifact_stores|jsonb|A mapping of artifactStore objects and their corresponding AWS Regions|
|version|integer|The version number of the pipeline|
