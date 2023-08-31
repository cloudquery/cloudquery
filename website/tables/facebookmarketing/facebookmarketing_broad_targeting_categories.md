# Table: facebookmarketing_broad_targeting_categories

This table shows data for Facebook Marketing Broad Targeting Categories.

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|category_description|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|parent_category|`utf8`|
|path|`list<item: utf8, nullable>`|
|size_lower_bound|`int64`|
|size_upper_bound|`int64`|
|source|`utf8`|
|type|`int64`|
|type_name|`utf8`|
|untranslated_name|`utf8`|
|untranslated_parent_name|`utf8`|