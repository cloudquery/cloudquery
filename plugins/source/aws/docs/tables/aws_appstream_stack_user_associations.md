# Table: aws_appstream_stack_user_associations

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UserStackAssociation.html

The composite primary key for this table is (**account_id**, **region**, **stack_name**, **user_name**, **authentication_type**).

## Relations
This table depends on [aws_appstream_stacks](aws_appstream_stacks.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|stack_name (PK)|String|
|user_name (PK)|String|
|authentication_type (PK)|String|
|send_email_notification|Bool|