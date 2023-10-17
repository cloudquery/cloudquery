This example configures a MongoDB destination, located at `localhost:27017`. The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).

```yaml copy
kind: destination
spec:
  name: "mongodb"
  registry: "github"
  path: "cloudquery/mongodb"
  version: "VERSION_DESTINATION_MONGODB"
  spec:
    connection_string: "mongodb://localhost:27017"
    database: "your_mongo_database_name"
    # Optional parameters:
    # batch_size: 10000 # 10K
    # batch_size_bytes: 4194304 # 4 MiB
```
