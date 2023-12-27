# Table: aws_docdb_pending_maintenance_actions

This table shows data for Amazon DocumentDB Pending Maintenance Actions.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_PendingMaintenanceAction.html

The composite primary key for this table is (**account_id**, **region**, **resource_identifier**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|pending_maintenance_action_details|`json`|
|resource_identifier (PK)|`utf8`|