This example configures a MongoDB destination, located at `localhost:27017`. The (top level) spec section is described in the [Destination Spec Reference](https://www.cloudquery.io/docs/cli/integrations/destinations#complete-destination-spec-reference).

```yaml copy
kind: destination
spec:
  name: "mongodb"
  path: "cloudquery/mongodb"
  registry: "cloudquery"
  version: "VERSION_DESTINATION_MONGODB"
  send_sync_summary: true
  spec:
    # required, a connection string in the format mongodb://localhost:27017
    connection_string: "${MONGODB_CONNECTION_STRING}"
    # required, the name of the database to sync to
    database: "${MONGODB_DATABASE_NAME}"
    # Optional parameters:
    # batch_size: 10000 # 10K
    # batch_size_bytes: 4194304 # 4 MiB
    # write_retry: # <- Opt-in retries on transient network errors. Safe with write_mode: overwrite, or with use_transactions on a replica set / sharded cluster.
    #   max_attempts: 5 # Total attempts per batch, including the first. Default 1 (retries disabled).
    #   max_backoff: "10s"
    #   use_transactions: false # Set true to wrap each retried batch in a MongoDB transaction (requires replica set / sharded cluster).
    # aws_credentials: # <- Use this to specify non-default role assumption parameters
    #   default: true # Use the default credentials chain
    #   local_profile: "mongodb-profile" # Use a local profile instead of the default one
    #   role_arn: "arn:aws:iam::123456789012:role/role_name" # Specify the role to assume
    #   external_id: "external_id" # Used when assuming a role
    #   role_session_name: "session_name" # Used when assuming a role

```
