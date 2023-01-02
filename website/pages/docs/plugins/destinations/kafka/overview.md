# Kafka Destination Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("destination", "kafka")}/>

This destination plugin lets you sync data from a CloudQuery source to Kafka in various formats such as CSV, JSON. Each table will be pushed to a separate topic.

## Example

This example configures connects to a local Kafka destination with no authentication and pushes messages in JSON format. Note that the plugin only supports `append` write-mode.

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml
kind: destination
spec:
  name: "kafka"
  path: "cloudquery/kafka"
  version: "VERSION_DESTINATION_KAFKA"
  write_mode: "append" # only supports 'append'

  spec:
    brokers: ["<broker-host>:<broker-port>"]
    format: "json"
```

## Plugin Spec

This is the (nested) plugin spec

- `brokers` (list(string)) (required)

  list of brokers to connect to.

- `format` (string) (required)

  Format of the output file. `json` and `csv` are supported.

- `sasl_username` (string) (optional)

  If connecting via SASL/PLAIN, the username to use.

- `sasl_password` (string) (optional)
  
  If connecting via SASL/PLAIN, the password to use.

- `verbose` (bool) (optional)

  If true, the plugin will log all underlying Kafka client messages to the log.