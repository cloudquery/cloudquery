# Amazon Kinesis Firehose Destination Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("destination", "azblob")}/>

This destination plugin lets you sync data from a CloudQuery source to remote Azure Blob Storage storage in various formats such as CSV, JSON.


## Authentication

TODO

## Example

TODO

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml
kind: destination
spec:
  name: "firehose"
  path: "cloudquery/firehose"
  version: "VERSION_DESTINATION_FIREHOSE"
  write_mode: "append" # this plugin only supports 'append' mode
  # TODO: Explain limitations of Batch Size
  # batch_size: 10000 # optional
  # batch_size_bytes: 5242880 # optional
  spec:
    storage_account: "cqdestinationfirehose"
    ARN
```
TODO
The Azure blob destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes). 

## Firehose Spec

