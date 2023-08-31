# Table: snyk_organization_provisions

This table shows data for Snyk Organization Provisions.

https://snyk.docs.apiary.io/#reference/organizations/provision-user/list-pending-user-provisions

The composite primary key for this table is (**organization_id**, **email**, **created**).

## Relations

This table depends on [snyk_organizations](snyk_organizations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id (PK)|`utf8`|
|email (PK)|`utf8`|
|role|`utf8`|
|role_public_id|`utf8`|
|created (PK)|`timestamp[us, tz=UTC]`|