# Table: snyk_organization_members

This table shows data for Snyk Organization Members.

https://snyk.docs.apiary.io/#reference/organizations/members-in-organization/list-members

The composite primary key for this table is (**organization_id**, **id**).

## Relations

This table depends on [snyk_organizations](snyk_organizations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id (PK)|`utf8`|
|id (PK)|`utf8`|
|username|`utf8`|
|name|`utf8`|
|email|`utf8`|
|role|`utf8`|