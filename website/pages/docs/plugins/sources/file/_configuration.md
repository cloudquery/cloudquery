```yaml copy
kind: source
spec:
  name: file
  path: cloudquery/file
  registry: cloudquery
  version: "VERSION_SOURCE_FILE"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  spec:
    files_dir: "/path/to/files-to-sync" # required. Path to the directory with files to sync
    # concurrency: 50 # optional. Number of files to sync in parallel. Default: 50
```
