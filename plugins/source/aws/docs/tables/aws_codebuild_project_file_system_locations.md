
# Table: aws_codebuild_project_file_system_locations
Information about a file system created by Amazon Elastic File System (EFS)
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_cq_id|uuid|Unique CloudQuery ID of aws_codebuild_projects table (FK)|
|identifier|text|The name used to access a file system created by Amazon EFS|
|location|text|A string that specifies the location of the file system created by Amazon EFS. Its format is efs-dns-name:/directory-path|
|mount_options|text|The mount options for a file system created by Amazon EFS|
|mount_point|text|The location in the container where you mount the file system.|
|type|text|The type of the file system|
