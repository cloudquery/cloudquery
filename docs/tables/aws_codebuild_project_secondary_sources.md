
# Table: aws_codebuild_project_secondary_sources
Information about the build input source code for the build project.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_cq_id|uuid|Unique CloudQuery ID of aws_codebuild_projects table (FK)|
|type|text|The type of repository that contains the source code to be built|
|auth_type|text|This data type is deprecated and is no longer accurate or used|
|auth_resource|text|The resource value that applies to the specified authorization type.|
|build_status_config_context|text|Specifies the context of the build status CodeBuild sends to the source provider|
|build_status_config_target_url|text|Specifies the target url of the build status CodeBuild sends to the source provider|
|buildspec|text|The buildspec file declaration to use for the builds in this build project|
|git_clone_depth|integer|Information about the Git clone depth for the build project.|
|git_submodules_config_fetch_submodules|boolean|Set to true to fetch Git submodules for your CodeBuild build project.|
|insecure_ssl|boolean|Enable this flag to ignore SSL warnings while connecting to the project source code.|
|location|text|Information about the location of the source code to be built|
|report_build_status|boolean|Set to true to report the status of a build's start and finish to your source provider|
|source_identifier|text|An identifier for this project source|
