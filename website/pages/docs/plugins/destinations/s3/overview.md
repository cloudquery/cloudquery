# S3 Destination Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("destination", "s3")}/>

This destination plugin lets you sync data from a CloudQuery source to remote S3 storage in various formats such as CSV, JSON.

This is useful in various use-cases, especially in data lakes where you can query the data direct from Athena or load it to various data warehouses such as BigQuery, RedShift, Snowflake and others.

## Authentication

Authentication is similar to AWS CLI. See also [aws source plugin](../../sources/aws/overview#authentication) for more information.

## Example

This example configures a CSV destination, to create CSV files in `s3://bucket_name/path/to/files`, with each table placed in its own directory. Note that the S3 plugin only supports `append` write-mode.

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml
kind: destination
spec:
  name: "s3"
  path: "cloudquery/s3"
  version: "VERSION_DESTINATION_S3"
  write_mode: "append" # s3 only supports 'append' mode
  # batch_size: 10000 # optional
  # batch_size_bytes: 5242880 # optional
  spec:
    bucket: "bucket_name"
    path: "path/to/files/{{TABLE}}/{{UUID}}.csv"
    format: "csv"
```

The S3 destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes).

## S3 Spec

This is the (nested) spec used by the CSV destination Plugin.

- `bucket` (string) (required)

  Bucket where to sync the files.

- `path` (string) (required)

  Path to where the files will be uploaded in the above bucket. The path supports two placeholder variables: `{{TABLE}}` will be replaced with the table name, and `{{UUID}}` with a random UUID to uniquely identify each file. 

- `format` (string) (required)

  Format of the output file. `json` and `csv` are supported.
