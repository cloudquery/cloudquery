# Table: aws_computeoptimizer_enrollment_statuses

This table shows data for Compute Optimizer Enrollment Statuses.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_GetEnrollmentStatus.html

The primary key for this table is **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|last_updated_timestamp|Timestamp|
|member_accounts_enrolled|Bool|
|number_of_member_accounts_opted_in|Int|
|status|String|
|status_reason|String|