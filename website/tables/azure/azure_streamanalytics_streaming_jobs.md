# Table: azure_streamanalytics_streaming_jobs

This table shows data for Azure Stream Analytics Streaming Jobs.

https://learn.microsoft.com/en-us/rest/api/streamanalytics/2020-03-01/streaming-jobs/list?tabs=HTTP#streamingjob

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|identity|`json`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|