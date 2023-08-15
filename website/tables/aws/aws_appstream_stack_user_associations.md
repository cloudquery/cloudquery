# Table: aws_appstream_stack_user_associations

This table shows data for Amazon AppStream Stack User Associations.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UserStackAssociation.html

The composite primary key for this table is (**account_id**, **region**, **stack_name**, **user_name**, **authentication_type**).

## Relations

This table depends on [aws_appstream_stacks](aws_appstream_stacks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|stack_name (PK)|`utf8`|
|user_name (PK)|`utf8`|
|authentication_type (PK)|`utf8`|
|send_email_notification|`bool`|