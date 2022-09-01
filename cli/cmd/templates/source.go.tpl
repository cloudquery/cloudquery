kind: source
spec:
  # Name of the plugin.
  name: "{{.Name}}"

  # Version of the plugin to use.
  version: "{{.Version}}"

  # Registry to use (one of "github", "local" or "grpc").
  registry: "{{.Registry}}"

  # Path to plugin. Required format depends on the registry.
  path: "{{.Path}}"

  # List of tables to sync.
  tables: ["*"]

  ## Tables to skip during sync. Optional.
  # skip_tables: []

  # Names of destination plugins to sync to.
  destinations: []

  ## Approximate cap on number of requests to perform concurrently. Optional.
  # max_goroutines: {{or .MaxGoRoutines 5 }}

  # Plugin-specific configuration.
  spec:
{{indent .Spec 4}}