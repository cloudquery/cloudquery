kind: "destination"
spec:
  # Name of the plugin.
  name: "{{.Name}}"

  # Version of the plugin to use.
  version: "{{.Version}}"

  # Registry to use (one of "github", "local" or "grpc").
  registry: "{{.Registry}}"

  # Path to plugin. Required format depends on the registry.
  path: "{{.Path}}"

  # Write mode (either "overwrite" or "append").
  write_mode: "{{.WriteMode}}"

  # Plugin-specific configuration.
  spec:
{{indent .Spec 4}}