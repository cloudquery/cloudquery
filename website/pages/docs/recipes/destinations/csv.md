# CSV Destination Plugin Recipes

Full spec options for CSV destination available [here](https://github.com/cloudquery/cloudquery/tree/main/plugins/destination/csv).


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
