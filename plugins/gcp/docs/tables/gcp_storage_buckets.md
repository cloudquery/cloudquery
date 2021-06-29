
# Table: gcp_storage_buckets
The Buckets resource represents a bucket in Cloud Storage
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|billing_requester_pays|boolean|When set to true, Requester Pays is enabled for this bucket|
|default_event_based_hold|boolean|The default value for event-based hold on newly created objects in this bucket Event-based hold is a way to retain objects indefinitely until an event occurs, signified by the hold's release After being released, such objects will be subject to bucket-level retention|
|encryption_default_kms_key_name|text|A Cloud KMS key that will be used to encrypt objects inserted into this bucket, if no encryption method is specified|
|etag|text|HTTP 11 Entity tag for the bucket|
|iam_configuration_bucket_policy_only_enabled|boolean|If set, access is controlled only by bucket-level or above IAM policies|
|iam_configuration_bucket_policy_only_locked_time|text|The deadline for changing iamConfigurationbucketPolicyOnlyenabled from true to false in RFC 3339 format iamConfigurationbucketPolicyOnlyenabled may be changed from true to false until the locked time, after which the field is immutable|
|iam_configuration_public_access_prevention|text|The bucket's Public Access Prevention configuration Currently, 'unspecified' and 'enforced' are supported|
|iam_configuration_uniform_bucket_level_access_enabled|boolean|If set, access is controlled only by bucket-level or above IAM policies|
|iam_configuration_uniform_bucket_level_access_locked_time|text|The deadline for changing iamConfigurationuniformBucketLevelAccessenabled from true to false in RFC 3339  format iamConfigurationuniformBucketLevelAccessenabled may be changed from true to false until the locked time, after which the field is immutable|
|resource_id|text|Original Id of the resource|
|kind|text|The kind of item this is For buckets, this is always storage#bucket|
|labels|jsonb|User-provided labels, in key/value pairs|
|location|text|The location of the bucket Object data for objects in the bucket resides in physical storage within this region Defaults to US See the developer's guide for the authoritative list|
|location_type|text|The type of the bucket location|
|logging_log_bucket|text|The destination bucket where the current bucket's logs should be placed|
|logging_log_object_prefix|text|A prefix for log object names|
|metageneration|bigint|The metadata generation of this bucket|
|name|text|The name of the bucket|
|owner_entity|text|The entity, in the form project-owner-projectId|
|owner_entity_id|text|The ID for the entity|
|project_number|bigint|The project number of the project the bucket belongs to|
|retention_policy_effective_time|text|Server-determined value that indicates the time from which policy was enforced and effective This value is in RFC 3339 format|
|retention_policy_is_locked|boolean|Once locked, an object retention policy cannot be modified|
|retention_policy_retention_period|bigint|The duration in seconds that objects need to be retained Retention duration must be greater than zero and less than 100 years Note that enforcement of retention periods less than a day is not guaranteed Such periods should only be used for testing purposes|
|satisfies_pzs|boolean|Reserved for future use|
|self_link|text|The URI of this bucket|
|storage_class|text|The bucket's default storage class, used whenever no storageClass is specified for a newly-created object This defines how objects in the bucket are stored and determines the SLA and the cost of storage Values include MULTI_REGIONAL, REGIONAL, STANDARD, NEARLINE, COLDLINE, ARCHIVE, and DURABLE_REDUCED_AVAILABILITY|
|time_created|text|The creation time of the bucket in RFC 3339 format|
|updated|text|The modification time of the bucket in RFC 3339 format|
|versioning_enabled|boolean|While set to true, versioning is fully enabled for this bucket|
|website_main_page_suffix|text|If the requested object path is missing, the service will ensure the path has a trailing '/', append this suffix, and attempt to retrieve the resulting object This allows the creation of indexhtml objects to represent directory pages|
|website_not_found_page|text|If the requested object path is missing, and any mainPageSuffix object is missing, if applicable, the service will return the named object from this bucket as the content for a 404 Not Found result|
|zone_affinity|text[]|The zone or zones from which the bucket is intended to use zonal quota Requests for data from outside the specified affinities are still allowed but won't be able to use zonal quota The zone or zones need to be within the bucket location otherwise the requests will fail with a 400 Bad Request response|
