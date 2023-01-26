# GCS (Google Cloud Storage) Destination Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("destination", "gcs")}/>

This destination plugin lets you sync data from a CloudQuery source to remote GCS (Google Cloud Storage) storage in various formats such as CSV, JSON.

This is useful in various use-cases, especially in data lakes where you can query the data direct from Athena or load it to various data warehouses such as BigQuery, RedShift, Snowflake and others.

## Example

This example configures a GCS destination, to create CSV files in `gcs://bucket_name/path/to/files`. Note that the GCS plugin only supports `append` write-mode.

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml
kind: destination
spec:
  name: "gcs"
  path: "cloudquery/gcs"
  version: "VERSION_DESTINATION_GCS"
  write_mode: "append" # gcs only supports 'append' mode
  # batch_size: 10000 # optional
  # batch_size_bytes: 5242880 # optional
  spec:
    bucket: "bucket_name"
    path: "path/to/files"
    format: "csv"
```

The GCS destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes).

## GCS Spec

This is the (nested) spec used by the CSV destination Plugin.

- `bucket` (string) (required)

  Bucket where to sync the files.

- `path` (string) (required)

  Path to where the files will be uploaded in the above bucket.

- `format` (string) (required)

  Format of the output file. `json` and `csv` are supported.
