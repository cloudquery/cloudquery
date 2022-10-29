# CloudQuery CSV Destination Plugin

This destination plugin let's you sync data from a CloudQuery source to a CSV files.

This is useful both in local environments but also in production environments where scalability, performance and cost are a requirements. You can use it to sync source in a distributed way and upload CSV to a remote storage (such as S3/GCS) and then load to datalakes such as BigQuery, Athena in batch mode.

## PostgreSQL Spec

This is the top level spec used by the PostgreSQL destination Plugin.

- `directory` (string) (optional, default to `./cq_csv_output`)

  Directory where all csv files will be written. a CSV file per table will be created.

