# Table: aws_auditmanager_assessments

This table shows data for AWS Audit Manager Assessments.

https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_Assessment.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|aws_account|`json`|
|framework|`json`|
|metadata|`json`|
|tags|`json`|