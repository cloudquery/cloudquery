# Table: snyk_groups

This table shows data for Snyk Groups.

https://snyk.docs.apiary.io/#reference/organizations/the-snyk-organization-for-a-request/list-all-the-organizations-a-user-belongs-to

	
This table lists all groups for the selected organizations. It uses the list organizations endpoint from the Snyk API.

The primary key for this table is **id**.

## Relations

The following tables depend on snyk_groups:
  - [snyk_group_members](snyk_group_members)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|