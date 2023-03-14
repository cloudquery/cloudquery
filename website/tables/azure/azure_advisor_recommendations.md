# Table: azure_advisor_recommendations

This table shows data for Azure Advisor Recommendations.

https://learn.microsoft.com/en-us/rest/api/advisor/recommendations/list?tabs=HTTP#resourcerecommendationbase

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|properties|JSON|
|id (PK)|String|
|name|String|
|type|String|