# Table: snyk_organization_provisions

This table shows data for Snyk Organization Provisions.

https://snyk.docs.apiary.io/#reference/organizations/provision-user/list-pending-user-provisions

The composite primary key for this table is (**organization_id**, **email**, **created**).

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
|email (PK)|String|
|role|String|
|role_public_id|String|
|created (PK)|String|