# CSV Destination Plugin Recipes

Full spec options for the CSV destination plugin are available [here](/docs/plugins/destinations/csv/overview#csv-spec).

## Basic

This is a basic configuration that will output all tables as CSV files to the specified directory.

```yaml copy
kind: destination
spec:
  name: csv
  path: cloudquery/csv
  version: "VERSION_DESTINATION_CSV"
  spec:
    directory: ./output # default to ./cq_csv_output
```
