
# Table: aws_codebuild_projects
Information about a build project.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the build project.|
|artifacts_type|text|The type of build output artifact|
|artifacts_artifact_identifier|text|An identifier for this artifact definition.|
|artifacts_bucket_owner_access|text|Specifies the bucket owner's access for objects that another account uploads to their Amazon S3 bucket|
|artifacts_encryption_disabled|boolean|Set to true if you do not want your output artifacts encrypted|
|artifacts_location|text|Information about the build output artifact location:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified|
|artifacts_name|text|Along with path and namespaceType, the pattern that CodeBuild uses to name and store the output artifact:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified|
|artifacts_namespace_type|text|Along with path and name, the pattern that CodeBuild uses to determine the name and location to store the output artifact:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified|
|artifacts_override_artifact_name|boolean|If this flag is set, a name specified in the buildspec file overrides the artifact name|
|artifacts_packaging|text|The type of build output artifact to create:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified|
|artifacts_path|text|Along with namespaceType and name, the pattern that CodeBuild uses to name and store the output artifact:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified|
|badge_enabled|boolean|Set this to true to generate a publicly accessible URL for your project's build badge.|
|badge_request_url|text|The publicly-accessible URL through which you can access the build badge for your project.|
|build_batch_config_batch_report_mode|text|Specifies how build status reports are sent to the source provider for the batch build|
|build_batch_config_combine_artifacts|boolean|Specifies if the build artifacts for the batch build should be combined into a single artifact location.|
|build_batch_config_restrictions_compute_types_allowed|text[]|An array of strings that specify the compute types that are allowed for the batch build|
|build_batch_config_restrictions_maximum_builds_allowed|integer|Specifies the maximum number of builds allowed.|
|build_batch_config_service_role|text|Specifies the service role ARN for the batch build project.|
|build_batch_config_timeout_in_mins|integer|Specifies the maximum amount of time, in minutes, that the batch build must be completed in.|
|cache_type|text|The type of cache used by the build project|
|cache_location|text|Information about the cache location:  * NO_CACHE or LOCAL: This value is ignored.  * S3: This is the S3 bucket name/prefix.|
|cache_modes|text[]|An array of strings that specify the local cache modes|
|concurrent_build_limit|integer|The maximum number of concurrent builds that are allowed for this project|
|created|timestamp without time zone|When the build project was created, expressed in Unix time format.|
|description|text|A description that makes the build project easy to identify.|
|encryption_key|text|The Key Management Service customer master key (CMK) to be used for encrypting the build output artifacts|
|environment_compute_type|text|Information about the compute resources the build project uses|
|environment_image|text|The image tag or image digest that identifies the Docker image to use for this build project|
|environment_type|text|The type of build environment to use for related builds.  * The environment type ARM_CONTAINER is available only in regions US East (N|
|environment_certificate|text|The ARN of the Amazon S3 bucket, path prefix, and object key that contains the PEM-encoded certificate for the build project|
|environment_image_pull_credentials_type|text|The type of credentials CodeBuild uses to pull images in your build|
|environment_privileged_mode|boolean|Enables running the Docker daemon inside a Docker container|
|environment_registry_credential|text|The Amazon Resource Name (ARN) or name of credentials created using Secrets Manager|
|environment_registry_credential_credential_provider|text|The service that created the credentials to access a private Docker registry. The valid value, SECRETS_MANAGER, is for Secrets Manager.|
|last_modified|timestamp without time zone|When the build project's settings were last modified, expressed in Unix time format.|
|logs_config_cloud_watch_logs_status|text|The current status of the logs in CloudWatch Logs for a build project|
|logs_config_cloud_watch_logs_group_name|text|The group name of the logs in CloudWatch Logs|
|logs_config_cloud_watch_logs_stream_name|text|The prefix of the stream name of the CloudWatch Logs|
|logs_config_s3_logs_status|text|The current status of the S3 build logs|
|logs_config_s3_logs_bucket_owner_access|text|Specifies the bucket owner's access for objects that another account uploads to their Amazon S3 bucket|
|logs_config_s3_logs_encryption_disabled|boolean|Set to true if you do not want your S3 build log output encrypted|
|logs_config_s3_logs_location|text|The ARN of an S3 bucket and the path prefix for S3 logs|
|name|text|The name of the build project.|
|project_visibility|text|Specifies the visibility of the project's builds|
|public_project_alias|text|Contains the project identifier used with the public build APIs.|
|queued_timeout_in_minutes|integer|The number of minutes a build is allowed to be queued before it times out.|
|resource_access_role|text|The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.|
|secondary_source_versions|jsonb|An array of ProjectSourceVersion objects|
|service_role|text|The ARN of the IAM role that enables CodeBuild to interact with dependent Amazon Web Services services on behalf of the Amazon Web Services account.|
|source_type|text|The type of repository that contains the source code to be built|
|source_auth_type|text|This data type is deprecated and is no longer accurate or used|
|source_auth_resource|text|The resource value that applies to the specified authorization type.|
|source_build_status_config_context|text|Specifies the context of the build status CodeBuild sends to the source provider|
|source_build_status_config_target_url|text|Specifies the target url of the build status CodeBuild sends to the source provider|
|source_buildspec|text|The buildspec file declaration to use for the builds in this build project|
|source_git_clone_depth|integer|Information about the Git clone depth for the build project.|
|source_git_submodules_config_fetch_submodules|boolean|Set to true to fetch Git submodules for your CodeBuild build project.|
|source_insecure_ssl|boolean|Enable this flag to ignore SSL warnings while connecting to the project source code.|
|source_location|text|Information about the location of the source code to be built|
|source_report_build_status|boolean|Set to true to report the status of a build's start and finish to your source provider|
|source_identifier|text|An identifier for this project source|
|source_version|text|A version of the build input to be built for this project|
|tags|jsonb|A list of tag key and value pairs associated with this build project|
|timeout_in_minutes|integer|How long, in minutes, from 5 to 480 (8 hours), for CodeBuild to wait before timing out any related build that did not get marked as completed|
|vpc_config_security_group_ids|text[]|A list of one or more security groups IDs in your Amazon VPC.|
|vpc_config_subnets|text[]|A list of one or more subnet IDs in your Amazon VPC.|
|vpc_config_vpc_id|text|The ID of the Amazon VPC.|
|webhook_branch_filter|text|A regular expression used to determine which repository branches are built when a webhook is triggered|
|webhook_build_type|text|Specifies the type of build this webhook will trigger.|
|webhook_filter_groups|jsonb|An array of arrays of WebhookFilter objects used to determine which webhooks are triggered|
|webhook_last_modified_secret|timestamp without time zone|A timestamp that indicates the last time a repository's secret token was modified.|
|webhook_payload_url|text|The CodeBuild endpoint where webhook events are sent.|
|webhook_secret|text|The secret token of the associated repository|
|webhook_url|text|The URL to the webhook.|
