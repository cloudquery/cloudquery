# CloudQuery CSV Destination Plugin

This destination plugin lets you sync data from a CloudQuery source to CSV files.

This is useful in local environments, but also in production environments where scalability, performance and cost are requirements. For example, this plugin can be used as part of a system that syncs sources across multiple virtual machines, uploads the CSVs to remote storage (such as S3 or GCS), and finally loads to datalakes such as BigQuery or Athena in batch mode.

## CSV Spec

This is the top level spec used by the CSV destination Plugin.

- `directory` (string) (optional, defaults to `./cq_csv_output`)

  Directory where all csv files will be written. A CSV file will be created per table.

