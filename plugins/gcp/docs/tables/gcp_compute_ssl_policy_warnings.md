
# Table: gcp_compute_ssl_policy_warnings

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|ssl_policy_id|uuid|Unique ID of gcp_compute_ssl_policies table (FK)|
|code|text|A warning code, if applicable For example, Compute Engine returns NO_RESULTS_ON_PAGE if there are no results in the response|
|data|jsonb|Metadata about this warning in key: value format|
|message|text|A human-readable description of the warning code|
