# Table: aws_organizations_organizational_units

https://docs.aws.amazon.com/organizations/latest/APIReference/API_OrganizationalUnit.html

The composite primary key for this table is (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|arn (PK)|String|
|id|String|
|name|String|