# Table: datadog_global_variables

This table shows data for Datadog Global Variables.

The composite primary key for this table is (**account_name**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|attributes|`json`|
|description|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|parse_test_options|`json`|
|parse_test_public_id|`utf8`|
|tags|`list<item: utf8, nullable>`|
|value|`json`|
|additional_properties|`json`|