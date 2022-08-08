
# Table: aws_elasticbeanstalk_application_versions
Describes the properties of an application version.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|application_name|text|The name of the application to which the application version belongs.|
|arn|text|The Amazon Resource Name (ARN) of the application version.|
|build_arn|text|Reference to the artifact from the AWS CodeBuild build.|
|date_created|timestamp without time zone|The creation date of the application version.|
|date_updated|timestamp without time zone|The last modified date of the application version.|
|description|text|The description of the application version.|
|source_location|text|The location of the source code, as a formatted string, depending on the value of SourceRepository  * For CodeCommit, the format is the repository name and commit ID, separated by a forward slash|
|source_repository|text|Location where the repository is stored.  * CodeCommit  * S3  This member is required.|
|source_type|text|The type of repository.  * Git  * Zip  This member is required.|
|source_bundle_s3_bucket|text|The Amazon S3 bucket where the data is located.|
|source_bundle_s3_key|text|The Amazon S3 key where the data is located.|
|status|text|The processing status of the application version|
|version_label|text|A unique identifier for the application version.|
