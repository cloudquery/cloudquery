# Table: github_traffic_referrers

https://docs.github.com/en/rest/metrics/traffic?apiVersion=2022-11-28#get-top-referral-sources

The composite primary key for this table is (**org**, **repository_id**, **referrer**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|repository_id (PK)|Int|
|referrer (PK)|String|
|count|Int|
|uniques|Int|