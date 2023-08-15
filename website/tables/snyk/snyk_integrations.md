# Table: snyk_integrations

This table shows data for Snyk Integrations.

https://snyk.docs.apiary.io/#reference/integrations/integrations/list

The composite primary key for this table is (**organization_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id (PK)|`utf8`|
|settings|`json`|
|credentials|`json`|
|id (PK)|`utf8`|
|type|`utf8`|