# Table: aws_computeoptimizer_enrollment_statuses

This table shows data for Compute Optimizer Enrollment Statuses.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_GetEnrollmentStatus.html

The primary key for this table is **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|last_updated_timestamp|`timestamp[us, tz=UTC]`|
|member_accounts_enrolled|`bool`|
|number_of_member_accounts_opted_in|`int64`|
|status|`utf8`|
|status_reason|`utf8`|