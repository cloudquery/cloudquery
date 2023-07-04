# Table: github_traffic_clones

This table shows data for Github Traffic Clones.

https://docs.github.com/en/rest/metrics/traffic?apiVersion=2022-11-28#get-repository-clones

The composite primary key for this table is (**org**, **repository_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|clones|`json`|
|count|`int64`|
|uniques|`int64`|