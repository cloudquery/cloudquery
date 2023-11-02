# Table: aws_codeartifact_repositories

This table shows data for AWS CodeArtifact Repositories.

https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_RepositoryDescription.html
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.

The composite primary key for this table is (**request_account_id**, **request_region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|tags|`json`|
|administrator_account|`utf8`|
|arn (PK)|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|domain_name|`utf8`|
|domain_owner|`utf8`|
|external_connections|`json`|
|name|`utf8`|
|upstreams|`json`|