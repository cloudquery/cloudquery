# Table: azure_security_sub_assessments

This table shows data for Azure Security Sub Assessments.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/sub-assessments/list?tabs=HTTP#securitysubassessment

The primary key for this table is **id**.

## Relations

This table depends on [azure_security_assessments](azure_security_assessments).

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