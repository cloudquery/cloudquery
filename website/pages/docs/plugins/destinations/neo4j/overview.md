# Neo4j Destination Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";
import { Callout } from 'nextra-theme-docs'

<Badge text={"Latest: " + getLatestVersion("destination", "neo4j")}/>

This destination plugin lets you sync data from a CloudQuery source to a Neo4j database.

Supported database (tested) versions (We use the [official Neo4j Go driver](https://github.com/neo4j/neo4j-go-driver#neo4j-and-bolt-protocol-versions)):

- Neo4j >= 4.4

As a side note graph databases can be quite useful for various networking use-cases, visualization, for read-teams, blue-teams and more.

## Configuration

### Example

This example configures a Neo4j destination, located at `bolt://localhost:7687`. The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml
kind: destination
spec:
  name: "neo4j"
  registry: "github"
  path: "cloudquery/neo4j"
  version: "VERSION_DESTINATION_NEO4J"
  # batch_size: 10000 # optional
  # batch_size_bytes: 5242880 # optional
  spec:
    connection_string: "bolt://localhost:7687"
    username: "${USERNAME}"
    password: "${PASSWORD}"
```

<Callout type="info">
Make sure you use environment variable expansion in production instead of committing the credentials to the configuration file directly.
</Callout>

The Neo4j destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes). 

### Plugin Spec

This is the (nested) spec used by the Neo4j destination Plugin.

- `connection_string` (string, required)

  Connection string to connect to the database. This can be a URL or a DSN, as per official [neo4j docs](https://neo4j.com/docs/browser-manual/current/operations/dbms-connection/#uri-scheme).

  - `"bolt://localhost:7687"`
  - `"neo4j://localhost:7687"`

- `username` (string, required)

  Username to connect to the database.

- `password` (string, required)

  Password to connect to the database.