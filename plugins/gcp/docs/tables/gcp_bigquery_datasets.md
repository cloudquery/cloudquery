
# Table: gcp_bigquery_datasets
dataset resources in the project
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|creation_time|bigint|The time when this dataset was created, in milliseconds since the epoch|
|default_encryption_configuration_kms_key_name|text|Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table The BigQuery Service Account associated with your project requires access to this encryption key|
|default_partition_expiration_ms|bigint|The default partition expiration for all partitioned tables in the dataset, in milliseconds Once this property is set, all newly-created partitioned tables in the dataset will have an expirationMs property in the timePartitioning settings set to this value, and changing the value will only affect new tables, not existing ones The storage in a partition will have an expiration time of its partition time plus this value Setting this property overrides the use of defaultTableExpirationMs for partitioned tables: only one of defaultTableExpirationMs and defaultPartitionExpirationMs will be used for any new partitioned table If you provide an explicit timePartitioningexpirationMs when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property|
|default_table_expiration_ms|bigint|The default lifetime of all tables in the dataset, in milliseconds The minimum value is 3600000 milliseconds (one hour) Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones When the expirationTime for a given table is reached, that table will be deleted automatically If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property|
|description|text|A user-friendly description of the dataset|
|etag|text|A hash of the resource|
|friendly_name|text|A descriptive name for the dataset|
|dataset_id|text|The fully-qualified unique name of the dataset in the format projectId:datasetId The dataset name without the project name is given in the datasetId field When creating a new dataset, leave this field blank, and instead specify the datasetId field|
|kind|text|The resource type|
|labels|jsonb|The labels associated with this dataset You can use these to organize and group your datasets You can set this property when inserting or updating a dataset See Creating and Updating Dataset Labels for more information|
|last_modified_time|bigint|The date when this dataset or any of its tables was last modified, in milliseconds since the epoch|
|location|text|The geographic location where the dataset should reside The default value is US See details at https://cloudgooglecom/bigquery/docs/locations|
|satisfies_pzs|boolean|Reserved for future use|
|self_link|text|A URL that can be used to access the resource again You can use this URL in Get or Update requests to the resource|
