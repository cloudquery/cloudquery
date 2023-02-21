# Table: aws_support_cases

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeCases.html

The composite primary key for this table is (**account_id**, **region**, **case_id**).

## Relations

The following tables depend on aws_support_cases:
  - [aws_support_case_communications](aws_support_case_communications.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|case_id (PK)|String|
|category_code|String|
|cc_email_addresses|StringArray|
|display_id|String|
|language|String|
|recent_communications|JSON|
|service_code|String|
|severity_code|String|
|status|String|
|subject|String|
|submitted_by|String|
|time_created|String|