# Table: snyk_organizations

This table shows data for Snyk Organizations.

https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to

The primary key for this table is **id**.

## Relations

The following tables depend on snyk_organizations:
  - [snyk_organization_members](snyk_organization_members)
  - [snyk_organization_provisions](snyk_organization_provisions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|group|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|slug|`utf8`|
|url|`utf8`|