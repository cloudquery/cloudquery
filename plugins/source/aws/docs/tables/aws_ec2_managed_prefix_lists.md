# Table: aws_ec2_managed_prefix_lists

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ManagedPrefixList.html. 
The 'request_account_id' and 'request_region' columns are added to show the account_id and region of where the request was made from.

The composite primary key for this table is (**request_account_id**, **request_region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|request_account_id (PK)|String|
|request_region (PK)|String|
|arn (PK)|String|
|tags|JSON|
|address_family|String|
|max_entries|Int|
|owner_id|String|
|prefix_list_arn|String|
|prefix_list_id|String|
|prefix_list_name|String|
|state|String|
|state_message|String|
|version|Int|