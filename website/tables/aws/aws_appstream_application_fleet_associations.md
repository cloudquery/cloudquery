# Table: aws_appstream_application_fleet_associations

This table shows data for Amazon AppStream Application Fleet Associations.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_ApplicationFleetAssociation.html

The composite primary key for this table is (**application_arn**, **fleet_name**).

## Relations

This table depends on [aws_appstream_applications](aws_appstream_applications).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|application_arn (PK)|`utf8`|
|fleet_name (PK)|`utf8`|