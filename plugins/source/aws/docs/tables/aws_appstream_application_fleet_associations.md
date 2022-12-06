# Table: aws_appstream_application_fleet_associations

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_ApplicationFleetAssociation.html

The composite primary key for this table is (**application_arn**, **fleet_name**).

## Relations
This table depends on [aws_appstream_applications](aws_appstream_applications.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|application_arn (PK)|String|
|fleet_name (PK)|String|