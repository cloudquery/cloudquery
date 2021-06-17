
# Table: gcp_storage_bucket_lifecycle_rules
A lifecycle management rule, which is made of an action to take and the condition(s) under which the action will be taken.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|bucket_id|uuid|Unique ID of gcp_storage_buckets table (FK)|
|action_storage_class|text|Target storage class Required iff the type of the action is SetStorageClass|
|action_type|text|Type of the action Currently, only Delete and SetStorageClass are supported|
|condition_age|bigint|Age of an object (in days) This condition is satisfied when an object reaches the specified age|
|condition_created_before|text|A date in RFC 3339 format with only the date part (for instance, "2013-01-15") This condition is satisfied when an object is created before midnight of the specified date in UTC|
|condition_custom_time_before|text|A date in RFC 3339 format with only the date part (for instance, "2013-01-15") This condition is satisfied when the custom time on an object is before this date in UTC|
|condition_days_since_custom_time|bigint|Number of days elapsed since the user-specified timestamp set on an object The condition is satisfied if the days elapsed is at least this number If no custom timestamp is specified on an object, the condition does not apply|
|condition_days_since_noncurrent_time|bigint|Number of days elapsed since the noncurrent timestamp of an object The condition is satisfied if the days elapsed is at least this number This condition is relevant only for versioned objects The value of the field must be a nonnegative integer If it's zero, the object version will become eligible for Lifecycle action as soon as it becomes noncurrent|
|condition_is_live|boolean|Relevant only for versioned objects If the value is true, this condition matches live objects; if the value is false, it matches archived objects|
|condition_matches_pattern|text|A regular expression that satisfies the RE2 syntax This condition is satisfied when the name of the object matches the RE2 pattern Note: This feature is currently in the "Early Access" launch stage and is only available to a whitelisted set of users; that means that this feature may be changed in backward-incompatible ways and that it is not guaranteed to be released|
|condition_matches_storage_class|text[]|Objects having any of the storage classes specified by this condition will be matched Values include MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, STANDARD, and DURABLE_REDUCED_AVAILABILITY|
|condition_noncurrent_time_before|text|A date in RFC 3339 format with only the date part (for instance, "2013-01-15") This condition is satisfied when the noncurrent time on an object is before this date in UTC This condition is relevant only for versioned objects|
|condition_num_newer_versions|bigint|Relevant only for versioned objects If the value is N, this condition is satisfied when there are at least N versions (including the live version) newer than this version of the object|
