# Table: aws_appstream_applications

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Application.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_appstream_applications:
  - [aws_appstream_application_fleet_associations](aws_appstream_application_fleet_associations.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|app_block_arn|String|
|created_time|Timestamp|
|description|String|
|display_name|String|
|enabled|Bool|
|icon_s3_location|JSON|
|icon_url|String|
|instance_families|StringArray|
|launch_parameters|String|
|launch_path|String|
|metadata|JSON|
|name|String|
|platforms|StringArray|
|working_directory|String|