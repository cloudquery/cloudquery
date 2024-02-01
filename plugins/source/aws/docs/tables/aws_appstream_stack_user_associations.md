# Table: aws_appstream_stack_user_associations

This table shows data for Amazon AppStream Stack User Associations.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UserStackAssociation.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **stack_name**, **user_name**, **authentication_type**).
## Relations

This table depends on [aws_appstream_stacks](aws_appstream_stacks.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|stack_name|`utf8`|
|user_name|`utf8`|
|authentication_type|`utf8`|
|send_email_notification|`bool`|