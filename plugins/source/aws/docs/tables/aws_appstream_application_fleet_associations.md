# Table: aws_appstream_application_fleet_associations

This table shows data for Amazon AppStream Application Fleet Associations.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_ApplicationFleetAssociation.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**application_arn**, **fleet_name**).
## Relations

This table depends on [aws_appstream_applications](aws_appstream_applications.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|application_arn|`utf8`|
|fleet_name|`utf8`|