# File Destination Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("destination", "file")}/>

This destination plugin lets you sync data from a CloudQuery source to local files in various formats such as CSV, JSON.

This is useful in local environments, but also in production environments where scalability, performance and cost are requirements. For example, this plugin can be used as part of a system that syncs sources across multiple virtual machines, uploads CSV files to a remote storage (such as S3 or GCS), and finally loads them to data lakes such as BigQuery or Athena in batch mode.

## Example

This example configures a CSV destination, to create CSV files in  `./cq_csv_output`. Note that the CSV plugin only supports `append` write-mode.

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml
kind: destination
spec:
  name: "file"
  path: "cloudquery/file"
  version: "VERSION_DESTINATION_FILE"
  write_mode: "append" # file only supports 'append' mode

  spec:
    directory: "./cq_csv_output"
    format: "csv"
```

## File Spec

This is the (nested) spec used by the CSV destination Plugin.

- `directory` (string) (required)

  Directory where all CSV files will be written. A CSV file will be created per table.

- `format` (string) (required)

  Format of the output file. `json` and `csv` are supported.

- `no_rotate` (bool) (optional)

  If set to true, the plugin will write to one file per table.
  Otherwise, for every batch a new file will be created with a different `.<UUID>` suffix.
