
# Table: github_hooks
Hook represents a GitHub (web and service) hook for a repository.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|org|text|The Github Organization of the resource.|
|created_at|timestamp without time zone||
|updated_at|timestamp without time zone||
|url|text||
|id|bigint||
|type|text||
|name|text||
|test_url|text||
|ping_url|text||
|last_response|jsonb||
|config|jsonb|Only the following fields are used when creating a hook. Config is required.|
|events|text[]||
|active|boolean||
