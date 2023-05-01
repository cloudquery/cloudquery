# Table: azure_security_sub_assessments

This table shows data for Azure Security Sub Assessments.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/sub-assessments/list?tabs=HTTP#securitysubassessment

The primary key for this table is **id**.

## Relations

This table depends on [azure_security_assessments](azure_security_assessments).

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