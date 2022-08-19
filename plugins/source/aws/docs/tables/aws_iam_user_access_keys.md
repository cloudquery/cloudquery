
# Table: aws_iam_user_access_keys

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|user_cq_id|uuid|Unique CloudQuery ID of aws_iam_users table (FK)|
|access_key_id|text|The ID for this access key.|
|create_date|timestamp without time zone|The date when the access key was created.|
|status|text|The status of the access key|
|user_name|text|The name of the IAM user that the key is associated with.|
|last_used_date|timestamp without time zone|The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the access key was most recently used. This field is null in the following situations:  * The user does not have an access key.  * An access key exists but has not been used since IAM began tracking this information.  * There is no sign-in data associated with the user.|
|last_used_region|text|The Amazon Web Services Region where this access key was most recently used|
|last_used_service_name|text|The name of the Amazon Web Services service with which this access key was most recently used|
