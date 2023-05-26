# Table: azure_security_assessments

This table shows data for Azure Security Assessments.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/assessments/list?tabs=HTTP#securityassessment

The primary key for this table is **id**.

## Relations

The following tables depend on azure_security_assessments:
  - [azure_security_sub_assessments](azure_security_sub_assessments)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|subscription_id|utf8|
|properties|json|
|id (PK)|utf8|
|name|utf8|
|type|utf8|