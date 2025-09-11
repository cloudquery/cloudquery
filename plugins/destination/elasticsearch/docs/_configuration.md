The following config will sync data to an Elasticsearch cluster running on `localhost:9200`:

```yaml copy
kind: destination
spec:
  name: elasticsearch
  path: cloudquery/elasticsearch
  registry: cloudquery
  version: "VERSION_DESTINATION_ELASTICSEARCH"
  write_mode: "overwrite-delete-stale"
  send_sync_summary: true
  spec:
    # Elastic Cloud configuration parameters
    cloud_id: "${ELASTICSEARCH_CLOUD_ID}"
    api_key: "${ELASTICSEARCH_API_KEY}"

    # Self-hosted Elasticsearch configuration parameters
    # addresses: ["http://localhost:9200"]
    # username: ""
    # password: ""
    # service_token: ""
    # certificate_fingerprint: ""
    # ca_cert: ""

    # Optional parameters
    # concurrency: 5 # default: number of CPUs
    # batch_size: 1000
    # batch_size_bytes: 5242880 # 5 MiB
```
