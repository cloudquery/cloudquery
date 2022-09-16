
# Table: aws_s3_bucket_replication_rules
Specifies which Amazon S3 objects to replicate and where to store the replicas.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|bucket_cq_id|uuid|Unique CloudQuery ID of aws_s3_buckets table (FK)|
|destination_bucket|text|The Amazon Resource Name (ARN) of the bucket where you want Amazon S3 to store the results.|
|destination_access_control_translation_owner|text|Specifies the replica ownership|
|destination_account|text|Destination bucket owner account ID|
|destination_encryption_configuration_replica_kms_key_id|text|Specifies the ID (Key ARN or Alias ARN) of the customer managed customer master key (CMK) stored in AWS Key Management Service (KMS) for the destination bucket. Amazon S3 uses this key to encrypt replica objects|
|destination_metrics_status|text|Specifies whether the replication metrics are enabled.|
|destination_metrics_event_threshold_minutes|integer|Contains an integer specifying time in minutes|
|destination_replication_time_status|text|Specifies whether the replication time is enabled.|
|destination_replication_time_minutes|integer|Contains an integer specifying time in minutes|
|destination_storage_class|text|The storage class to use when replicating objects, such as S3 Standard or reduced redundancy|
|status|text|Specifies whether the rule is enabled.|
|delete_marker_replication_status|text|Indicates whether to replicate delete markers|
|existing_object_replication_status|text||
|filter|jsonb|A filter that identifies the subset of objects to which the replication rule applies|
|id|text|A unique identifier for the rule|
|prefix|text|An object key name prefix that identifies the object or objects to which the rule applies|
|priority|integer|The priority indicates which rule has precedence whenever two or more replication rules conflict|
|source_replica_modifications_status|text|Specifies whether Amazon S3 replicates modifications on replicas.|
|source_sse_kms_encrypted_objects_status|text|Specifies whether Amazon S3 replicates objects created with server-side encryption using a customer master key (CMK) stored in AWS Key Management Service.|
