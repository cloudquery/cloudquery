# Table: aws_codecommit_repositories

This table shows data for AWS CodeCommit Repositories.

https://docs.aws.amazon.com/codecommit/latest/APIReference/API_RepositoryMetadata.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|clone_url_http|`utf8`|
|clone_url_ssh|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|default_branch|`utf8`|
|last_modified_date|`timestamp[us, tz=UTC]`|
|repository_description|`utf8`|
|repository_id|`utf8`|
|repository_name|`utf8`|