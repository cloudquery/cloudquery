# Table: test_paid_table

This is a premium table. To sync this table you must be logged in via `cloudquery login` or you must use a valid API Key which can be generated at `cloud.cloudquery.io`

Test Paid table

The composite primary key for this table is (**resource_id**, **client_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|resource_id (PK)|`int64`|
|column2|`utf8`|
|client_id (PK)|`int64`|