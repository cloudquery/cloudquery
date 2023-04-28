# Table: snyk_organization_members

This table shows data for Snyk Organization Members.

https://snyk.docs.apiary.io/#reference/organizations/members-in-organization/list-members

The composite primary key for this table is (**organization_id**, **id**).

## Relations

This table depends on [snyk_organizations](snyk_organizations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|organization_id (PK)|String|
|id (PK)|String|
|username|String|
|name|String|
|email|String|
|role|String|