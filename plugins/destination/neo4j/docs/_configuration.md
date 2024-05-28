This example configures a Neo4j destination, located at `bolt://localhost:7687`. The username and password are stored in environment variables.

```yaml copy
kind: destination
spec:
  name: "neo4j"
  path: "cloudquery/neo4j"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_NEO4J"
  # Learn more about the configuration options at https://cql.ink/neo4j_destination
  spec:
    connection_string: "${NEO4J_CONNECTION_STRING}"
    username: "${NEO4J_USERNAME}"
    password: "${NEO4J_PASSWORD}"
    # Optional parameters:
    # batch_size: 1000 # 1K entries
    # batch_size_bytes: 4194304 # 4 MiB
```
