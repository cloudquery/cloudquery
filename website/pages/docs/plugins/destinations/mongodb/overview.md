# MongoDB Destination Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";
import { Callout } from 'nextra-theme-docs'

<Badge text={"Latest: " + getLatestVersion("destination", "mongodb")}/>

This destination plugin lets you sync data from any CloudQuery source to a MongoDB database.

Supported database versions:

- MongoDB >= 3.6 (The same minimum version supported by the official [Go driver](https://github.com/mongodb/mongo-go-driver))

## Configuration

### Example

This example configures a MongoDB destination, located at `localhost:27017`. The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml
kind: destination
spec:
  name: "mongodb"
  registry: "github"
  path: "cloudquery/mongodb"
  version: "VERSION_DESTINATION_MONGODB"
  spec:
    connection_string: "mongodb://localhost:27017"
    database: "your_mongo_database_name"
```

<Callout type="info">
Make sure to use [environment variable substitution](/docs/advanced-topics/environment-variable-substitution) in production instead of committing the credentials to the configuration file directly.
</Callout>

### MongoDB Spec

This is the (nested) spec used by the MongoDB destination Plugin.

- `connection_string` (string, required)

  MongoDB URI as described in the official MongoDB [documentation](https://www.mongodb.com/docs/manual/reference/connection-string/)

- `database` (string, required)

  Required database to sync the data to




