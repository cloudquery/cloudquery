The following config will sync data to a Meilisearch instance running on `localhost:7700`:

```yaml copy
kind: destination
spec:
  name: meilisearch
  path: cloudquery/meilisearch
  registry: cloudquery
  version: "VERSION_DESTINATION_MEILISEARCH"
  write_mode: "overwrite"
  # Learn more about the configuration options at https://cql.ink/meilisearch_destination
  spec:
    # meilisearch plugin spec
    host: "${MEILISEARCH_HOST}"
    api_key: "${MEILISEARCH_API_KEY}"
    # Optional parameters:
    # timeout: 5m
    # ca_cert: "<YOUR_MEILISEARCH_CA_CERT>"
    # batch_size: 1000 # 1K entries
    # batch_size_bytes: 4194304 # 4 MiB
    # batch_timeout: 20s
```
