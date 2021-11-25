
# Table: aws_codebuild_project_secondary_artifacts
Information about the build output artifacts for the build project.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_cq_id|uuid|Unique CloudQuery ID of aws_codebuild_projects table (FK)|
|type|text|The type of build output artifact|
|artifact_identifier|text|An identifier for this artifact definition.|
|bucket_owner_access|text|Specifies the bucket owner's access for objects that another account uploads to their Amazon S3 bucket|
|encryption_disabled|boolean|Set to true if you do not want your output artifacts encrypted|
|location|text|Information about the build output artifact location:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified|
|name|text|Along with path and namespaceType, the pattern that CodeBuild uses to name and store the output artifact:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified|
|namespace_type|text|Along with path and name, the pattern that CodeBuild uses to determine the name and location to store the output artifact:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified|
|override_artifact_name|boolean|If this flag is set, a name specified in the buildspec file overrides the artifact name|
|packaging|text|The type of build output artifact to create:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified|
|path|text|Along with namespaceType and name, the pattern that CodeBuild uses to name and store the output artifact:  * If type is set to CODEPIPELINE, CodePipeline ignores this value if specified|
