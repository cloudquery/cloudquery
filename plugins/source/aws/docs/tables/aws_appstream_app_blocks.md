# Table: aws_appstream_app_blocks

This table shows data for Amazon AppStream App Blocks.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_AppBlock.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|name|`utf8`|
|app_block_errors|`json`|
|created_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|display_name|`utf8`|
|packaging_type|`utf8`|
|post_setup_script_details|`json`|
|setup_script_details|`json`|
|source_s3_location|`json`|
|state|`utf8`|