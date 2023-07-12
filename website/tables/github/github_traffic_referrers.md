# Table: github_traffic_referrers

This table shows data for Github Traffic Referrers.

https://docs.github.com/en/rest/metrics/traffic?apiVersion=2022-11-28#get-top-referral-sources

The composite primary key for this table is (**org**, **repository_id**, **referrer**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|referrer (PK)|`utf8`|
|count|`int64`|
|uniques|`int64`|