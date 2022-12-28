# S3 Destination Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("destination", "s3")}/>

This destination plugin lets you sync data from a CloudQuery source to remote S3 storage in various formats such as CSV, JSON.

This is useful in various use-cases, especially in datalakes where you can query the data direct from Athena or load it to various data warehouses such as BigQuery, RedShift, Snowflake and others.

## Example

This example configures a CSV destination, to create CSV files in `s3://bucket_name/path/to/files`. Note that the CSV plugin only supports `append` write-mode.

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml
kind: destination
spec:
  name: "s3"
  path: "cloudquery/s3"
  version: "VERSION_DESTINATION_S3"
  write_mode: "append" # s3 only supports 'append' mode

  spec:
    bucket: "bucket_name"
    key: "path/to/files"
    format: "csv"
```

## S3 Spec

This is the (nested) spec used by the CSV destination Plugin.

- `bucket` (string) (required)

  Bucket where to sync the files.

- `key` (string) (required)

  Path to where the files will be uploaded in the above bucket.

- `format` (string) (required)

  Format of the output file. `json` and `csv` are supported.
