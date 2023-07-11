# Table: azure_advisor_recommendations

This table shows data for Azure Advisor Recommendations.

https://learn.microsoft.com/en-us/rest/api/advisor/recommendations/list?tabs=HTTP#resourcerecommendationbase

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|