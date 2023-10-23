```yaml copy
kind: source
spec:
  name: file
  path: cloudquery/file # Buy from here: https://cloudquery.io/integrations/file
  registry: cloudquery
  version: "PREMIUM"
  tables: ["*"]
  destinations: ["DESTINATION_NAME"]

  spec:
    files_dir: "/path/to/files-to-sync" # required. Path to the directory with files to sync
    # concurrency: 50 # optional. Number of files to sync in parallel. Default: 50
```
