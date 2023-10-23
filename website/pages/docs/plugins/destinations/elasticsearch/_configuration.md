The following config will sync data to an Elasticsearch cluster running on `localhost:9200`:

```yaml copy
kind: destination
spec:
  name: elasticsearch
  path: cloudquery/elasticsearch
  version: "VERSION_DESTINATION_ELASTICSEARCH"
  write_mode: "overwrite-delete-stale"
  spec:
    # Optional parameters
    # addresses: ["http://localhost:9200"]
    # username: ""
    # password: ""
    # cloud_id: ""
    # api_key: ""
    # service_token: ""
    # certificate_fingerprint: ""
    # ca_cert: ""
    # concurrency: 5 # default: number of CPUs
    # batch_size: 1000
    # batch_size_bytes: 5242880 # 5 MiB
```