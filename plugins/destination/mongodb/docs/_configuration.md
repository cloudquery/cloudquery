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
    # aws_credentials: # <- Use this to specify non-default role assumption parameters
    #   default: true # Use the default credentials chain
    #   local_profile: "mongodb-profile" # Use a local profile instead of the default one
    #   role_arn: "arn:aws:iam::123456789012:role/role_name" # Specify the role to assume
    #   external_id: "external_id" # Used when assuming a role
    #   role_session_name: "session_name" # Used when assuming a role
    # oidc: # <- Use this for MONGODB-OIDC Workload Identity Federation (overrides connection_string credentials)
    #   environment: "gcp" # One of `gcp`, `azure` or `k8s`
    #   token_resource: "${MONGODB_OIDC_TOKEN_RESOURCE}" # Audience configured on the Atlas deployment; required for `gcp` and `azure`

```

When `oidc` is set, the plugin authenticates with the `MONGODB-OIDC` mechanism using the driver's built-in provider for the given `environment`. This requires [Workload Identity Federation](https://www.mongodb.com/docs/atlas/workload-oidc/) to be configured on the Atlas side (OIDC identity provider, matching audience, and a database user mapped to the federated identity), and the workload to be running with a service account the chosen provider can read (for example, a GCP service account on Cloud Run, GCE or GKE).
