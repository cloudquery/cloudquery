```yaml copy
kind: source
spec:
  # Source spec section
  name: "gcp"
  path: "cloudquery/gcp"
  version: "VERSION_SOURCE_GCP"
  tables: ["gcp_storage_buckets"]
  destinations: ["DESTINATION_NAME"]
  spec:
    # GCP Spec
    project_ids: ["my-project"]
```
