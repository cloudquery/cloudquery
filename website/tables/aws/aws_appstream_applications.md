# Table: aws_appstream_applications

This table shows data for Amazon AppStream Applications.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Application.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_appstream_applications:
  - [aws_appstream_application_fleet_associations](aws_appstream_application_fleet_associations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|app_block_arn|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|display_name|`utf8`|
|enabled|`bool`|
|icon_s3_location|`json`|
|icon_url|`utf8`|
|instance_families|`list<item: utf8, nullable>`|
|launch_parameters|`utf8`|
|launch_path|`utf8`|
|metadata|`json`|
|name|`utf8`|
|platforms|`list<item: utf8, nullable>`|
|working_directory|`utf8`|