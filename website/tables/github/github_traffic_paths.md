# Table: github_traffic_paths

This table shows data for Github Traffic Paths.

https://docs.github.com/en/rest/metrics/traffic?apiVersion=2022-11-28#get-top-referral-paths

The composite primary key for this table is (**org**, **repository_id**, **path**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|path (PK)|`utf8`|
|title|`utf8`|
|count|`int64`|
|uniques|`int64`|