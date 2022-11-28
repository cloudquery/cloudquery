# CloudQuery CSV Destination Plugin

This destination plugin lets you sync data from a CloudQuery source to CSV files.

This is useful in local environments, but also in production environments where scalability, performance and cost are requirements. For example, this plugin can be used as part of a system that syncs sources across multiple virtual machines, uploads CSV files to a remote storage (such as S3 or GCS), and finally loads them to data lakes such as BigQuery or Athena in batch mode.

## Example

This example configures a CSV destination, to create CSV files in  `./cq_csv_output`. Note that the CSV plugin only supports `append` write-mode.

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml
kind: destination
spec:
  name: "csv"
  path: "cloudquery/csv"
  version: "VERSION_DESTINATION_CSV"
  write_mode: "append" # CSV only supports 'append' mode

  spec:
    directory: './cq_csv_output'
```

## CSV Spec

This is the (nested) spec used by the CSV destination Plugin.

- `directory` (string) (optional, defaults to `./cq_csv_output`)

  Directory where all CSV files will be written. A CSV file will be created per table.
