
# Table: gcp_bigquery_dataset_table_view_user_defined_function_resources
This is used for defining User Defined Function (UDF) resources only when using legacy SQL Users of Standard SQL should leverage either DDL (eg CREATE [TEMPORARY] FUNCTION  ) or the Routines API to define UDF resources For additional information on migrating, see: https://cloudgoogle
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|dataset_table_id|uuid|Unique ID of gcp_bigquery_dataset_tables table (FK)|
|inline_code|text|[Pick one] An inline resource that contains code for a user-defined function (UDF) Providing a inline code resource is equivalent to providing a URI for a file containing the same code|
|resource_uri|text|[Pick one] A code resource to load from a Google Cloud Storage URI (gs://bucket/path)|
