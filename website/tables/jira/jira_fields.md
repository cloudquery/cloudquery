# Table: jira_fields

This table shows data for Jira Fields.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`utf8`|
|key|`utf8`|
|name|`utf8`|
|custom|`bool`|
|navigable|`bool`|
|searchable|`bool`|
|clause_names|`list<item: utf8, nullable>`|
|schema|`json`|