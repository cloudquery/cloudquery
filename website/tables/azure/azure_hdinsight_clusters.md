# Table: azure_hdinsight_clusters

This table shows data for Azure HDInsight Clusters.

https://learn.microsoft.com/en-us/rest/api/hdinsight/2021-06-01/clusters/list?tabs=HTTP#cluster

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|etag|`utf8`|
|identity|`json`|
|properties|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|