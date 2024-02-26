This example configures a MongoDB destination, located at `localhost:27017`. The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml copy
kind: destination
spec:
  name: "mongodb"
  path: "cloudquery/mongodb"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_MONGODB"
  spec:
    # required, a connection string in the format mongodb://localhost:27017
    connection_string: "${MONGODB_CONNECTION_STRING}"
    # required, the name of the database to sync to
    database: "${MONGODB_DATABASE_NAME}"
    # Optional parameters:
    # batch_size: 10000 # 10K
    # batch_size_bytes: 4194304 # 4 MiB
```
