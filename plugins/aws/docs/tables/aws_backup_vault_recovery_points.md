
# Table: aws_backup_vault_recovery_points
The recovery points stored in a backup vault.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|vault_cq_id|uuid|Unique CloudQuery ID of aws_backup_vault table (FK)|
|backup_size|bigint|The size, in bytes, of a backup.|
|calculated_delete_at|timestamp without time zone|A timestamp that specifies when to delete a recovery point.|
|calculated_move_to_cold_storage_at|timestamp without time zone|A timestamp that specifies when to transition a recovery point to cold storage.|
|completion_date|timestamp without time zone|The date and time a job to restore a recovery point is completed, in Unix format and Coordinated Universal Time (UTC).|
|created_by|jsonb|Contains identifying information about the creation of a recovery point.|
|creation_date|timestamp without time zone|The date and time a recovery point is created, in Unix format and Coordinated Universal Time (UTC).|
|encryption_key_arn|text|The server-side encryption key that is used to protect your backups.|
|iam_role_arn|text|Specifies the IAM role ARN used to create the target recovery point.|
|is_encrypted|boolean|Describes if the recovery point is encrypted.|
|last_restore_time|timestamp without time zone|The date and time a recovery point was last restored, in Unix format and Coordinated Universal Time (UTC).|
|delete_after|bigint|Specifies the number of days after creation that a recovery point is deleted.|
|move_to_cold_storage_after|bigint|Specifies the number of days after creation that a recovery point is moved to cold storage.|
|arn|text|An Amazon Resource Name (ARN) that uniquely identifies a recovery point.|
|resource_arn|text|An ARN that uniquely identifies a resource (saved as a recovery point).|
|resource_type|text|The type of Amazon Web Services resource saved as a recovery point.|
|source_backup_vault_arn|text|The backup vault where the recovery point was originally copied from.|
|status|text|A status code specifying the state of the recovery point.|
|status_message|text|A message explaining the reason of the recovery point deletion failure.|
|tags|jsonb|Resource tags|
