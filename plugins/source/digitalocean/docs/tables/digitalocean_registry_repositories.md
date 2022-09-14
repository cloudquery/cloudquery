
# Table: digitalocean_registry_repositories
Repository represents a repository
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|registry_cq_id|uuid|Unique CloudQuery ID of digitalocean_registry table (FK)|
|registry_name|text|The name of the container registry.|
|name|text|The name of the repository.|
|latest_tag_registry_name|text|The name of the container registry.|
|latest_tag_repository|text|The name of the repository.|
|latest_tag|text|The name of the tag.|
|latest_tag_manifest_digest|text|The digest of the manifest associated with the tag.|
|latest_tag_compressed_size_bytes|bigint|The compressed size of the tag in bytes.|
|latest_tag_size_bytes|bigint|The uncompressed size of the tag in bytes (this size is calculated asynchronously so it may not be immediately available).|
|latest_tag_updated_at|timestamp without time zone|The time the tag was last updated.|
|tag_count|bigint|The number of tags in the repository.|
