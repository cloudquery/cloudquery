# Configuration

<details>
<summary>Relevant source files</summary>

The following files were used as context for generating this wiki page:

- [cli/cmd/testdata/cloudflare-tables-with-spec-filter.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/cloudflare-tables-with-spec-filter.yml)
- [cli/cmd/testdata/different-backend-from-destination.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/different-backend-from-destination.yml)
- [cli/cmd/testdata/multiple-destinations.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/multiple-destinations.yml)
- [cli/cmd/testdata/multiple-sources-destinations.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/multiple-sources-destinations.yml)
- [cli/cmd/testdata/multiple-sources.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/multiple-sources.yml)
- [cli/cmd/testdata/sync-missing-path-error.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/sync-missing-path-error.yml)
- [cli/cmd/testdata/sync-success-sourcev1-destv0.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/sync-success-sourcev1-destv0.yml)
- [cli/cmd/testdata/test-connection-bad-connection.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/test-connection-bad-connection.yml)
- [cli/cmd/testdata/validate-config-error.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/validate-config-error.yml)
- [cli/cmd/testdata/with-destination-summary-with-sync-group-id-and-shard.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/with-destination-summary-with-sync-group-id-and-shard.yml)
- [cli/cmd/testdata/with-destination-summary.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/with-destination-summary.yml)

</details>



This page documents how to configure CloudQuery through YAML files. CloudQuery uses a declarative configuration approach to define source plugins that extract data, destination plugins that load data, and their respective settings.

## Configuration File Structure

CloudQuery configuration files use YAML format and consist of one or more configuration blocks separated by `---`. Each block defines either a source or destination plugin with its configuration.

```mermaid
graph TD
    subgraph "Configuration File"
        ConfigFile["YAML Configuration"]
        SourceBlocks["Source Blocks"]
        DestBlocks["Destination Blocks"]
        
        ConfigFile --> SourceBlocks
        ConfigFile --> DestBlocks
        
        subgraph "Source Block"
            SourceKind["kind: source"]
            SourceSpec["spec:"]
            SourceName["name"]
            SourcePath["path"]
            SourceVersion["version"]
            SourceDestinations["destinations"]
            SourceTables["tables"]
            SourcePluginSpec["spec"]
            
            SourceKind --> SourceSpec
            SourceSpec --> SourceName
            SourceSpec --> SourcePath
            SourceSpec --> SourceVersion
            SourceSpec --> SourceDestinations
            SourceSpec --> SourceTables
            SourceSpec --> SourcePluginSpec
        end
        
        subgraph "Destination Block"
            DestKind["kind: destination"]
            DestSpec["spec:"]
            DestName["name"]
            DestPath["path"]
            DestVersion["version"]
            DestPluginSpec["spec"]
            
            DestKind --> DestSpec
            DestSpec --> DestName
            DestSpec --> DestPath
            DestSpec --> DestVersion
            DestSpec --> DestPluginSpec
        end
    end
```

Sources: [cli/cmd/testdata/multiple-sources-destinations.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/multiple-sources-destinations.yml), [cli/cmd/testdata/validate-config-error.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/validate-config-error.yml)

### Basic Configuration Structure

Each configuration block has two main parts:
- `kind`: Specifies whether the block defines a `source` or `destination`
- `spec`: Contains the configuration details for the plugin

## Source Configuration

Source plugins define data sources from which CloudQuery extracts information.

### Required Fields

| Field | Description |
|-------|-------------|
| `name` | Unique identifier for the source |
| `path` | Plugin path (format: `registry/plugin`) |
| `version` | Plugin version (semantic versioning) |
| `destinations` | List of destination names to send data to |
| `tables` | List of tables to sync (use `["*"]` for all tables) |

### Optional Fields

| Field | Description |
|-------|-------------|
| `registry` | Plugin registry (`cloudquery`, `github`, etc.) |
| `backend_options` | Configuration for state backend storage |

### Plugin-Specific Configuration

Each source plugin has its own configuration specified in the nested `spec` field.

```yaml
kind: source
spec:
  name: cloudflare
  path: cloudquery/cloudflare
  version: "v10.0.0" 
  destinations: ["postgresql"]
  tables: ["*"]
  spec:
    api_token: "your-api-token"
```

Sources: [cli/cmd/testdata/cloudflare-tables-with-spec-filter.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/cloudflare-tables-with-spec-filter.yml), [cli/cmd/testdata/test-connection-bad-connection.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/test-connection-bad-connection.yml)

## Destination Configuration

Destination plugins define where CloudQuery loads the extracted data.

### Required Fields

| Field | Description |
|-------|-------------|
| `name` | Unique identifier for the destination |
| `path` | Plugin path (format: `registry/plugin`) |
| `version` | Plugin version (semantic versioning) |

### Optional Fields

| Field | Description |
|-------|-------------|
| `registry` | Plugin registry (`cloudquery`, `github`, etc.) |
| `send_sync_summary` | Boolean to enable sync summary sending |
| `sync_group_id` | Identifier for grouping syncs |
| `write_mode` | How data is written (e.g., `append`) |

### Plugin-Specific Configuration

Each destination plugin has its own configuration specified in the nested `spec` field.

```yaml
kind: destination
spec:
  name: "postgresql"
  path: "cloudquery/postgresql"
  version: "v8.8.2"
  spec:
    connection_string: "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable"
```

Sources: [cli/cmd/testdata/validate-config-error.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/validate-config-error.yml), [cli/cmd/testdata/with-destination-summary.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/with-destination-summary.yml)

## Configuration Relationships

The following diagram illustrates how different components in CloudQuery configuration relate to each other:

```mermaid
graph TD
    ConfigFile["Configuration File"]
    SourcePlugin1["Source Plugin 1"]
    SourcePlugin2["Source Plugin 2"]
    DestPlugin1["Destination Plugin 1"]
    DestPlugin2["Destination Plugin 2"]
    
    ConfigFile --> SourcePlugin1
    ConfigFile --> SourcePlugin2
    ConfigFile --> DestPlugin1
    ConfigFile --> DestPlugin2
    
    SourcePlugin1 --> |"destinations: [dest1]"| DestPlugin1
    SourcePlugin2 --> |"destinations: [dest1, dest2]"| DestPlugin1
    SourcePlugin2 --> |"destinations: [dest1, dest2]"| DestPlugin2
    
    subgraph "Data Flow"
        SourcePlugin1 --> Extract["Extract Data"]
        SourcePlugin2 --> Extract
        Extract --> Transform["Transform Data"]
        Transform --> Load["Load Data"]
        Load --> DestPlugin1
        Load --> DestPlugin2
    end
```

Sources: [cli/cmd/testdata/multiple-sources-destinations.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/multiple-sources-destinations.yml), [cli/cmd/testdata/multiple-destinations.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/multiple-destinations.yml)

## Advanced Configuration Examples

### Multiple Sources and Destinations

CloudQuery supports configuring multiple sources and destinations in a single configuration file:

```yaml
kind: "source"
spec:
  name: "test-1"
  path: "cloudquery/test"
  destinations: [test-1]
  version: "v4.5.1"
  tables: ["*"]
---
kind: "source"
spec:
  name: "test-2"
  path: "cloudquery/test"
  destinations: [test-2]
  version: "v4.5.1"
  tables: ["*"]
---
kind: "destination"
spec:
  name: "test-1"
  path: "cloudquery/test"
  version: "v2.5.1"
---
kind: "destination"
spec:
  name: "test-2"
  path: "cloudquery/test"
  version: "v2.5.1"
```

Sources: [cli/cmd/testdata/multiple-sources-destinations.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/multiple-sources-destinations.yml)

### Multiple Sources to Single Destination

You can configure multiple sources to send data to the same destination:

```yaml
kind: "source"
spec:
  name: "test"
  path: "cloudquery/test"
  destinations: [test]
  version: "v4.5.1"
  tables: ["*"]
---
kind: "source"
spec:
  name: "test2"
  path: "cloudquery/test"
  destinations: [test]
  version: "v4.5.1"
  tables: ["*"]
---
kind: "destination"
spec:
  name: "test"
  path: "cloudquery/test"
  version: "v2.5.1"
```

Sources: [cli/cmd/testdata/multiple-sources.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/multiple-sources.yml)

### Single Source to Multiple Destinations

A single source can send data to multiple destinations:

```yaml
kind: "source"
spec:
  name: "test"
  path: "cloudquery/test"
  destinations: ["test1", "test2"]
  version: "v2.0.3"
  tables: ["*"]
---
kind: "destination"
spec:
  name: "test1"
  path: "cloudquery/test"
  version: "v2.1.0"
---
kind: "destination"
spec:
  name: "test2"
  path: "cloudquery/test"
  version: "v2.5.1"
```

Sources: [cli/cmd/testdata/multiple-destinations.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/multiple-destinations.yml)

### Backend Configuration

You can configure a separate backend for state storage:

```yaml
kind: "source"
spec:
  name: "test"
  path: "cloudquery/test"
  destinations: ["test1"]
  backend_options:
    table_name: "test_backend"
    connection: "@@plugins.test2.connection"
  version: "v4.5.1"
  tables: ["*"]
---
kind: "destination"
spec:
  name: "test1"
  path: "cloudquery/test"
  version: "v2.5.1"
---
kind: "destination"
spec:
  name: "test2"
  path: "cloudquery/test"
  version: "v2.5.1"
```

Sources: [cli/cmd/testdata/different-backend-from-destination.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/different-backend-from-destination.yml)

## Configuration Processing Flow

The following diagram illustrates how CloudQuery processes configuration files:

```mermaid
graph TD
    ConfigFile["YAML Configuration File"]
    ParseYAML["Parse YAML"]
    ValidateConfig["Validate Configuration"]
    InitPlugins["Initialize Plugins"]
    FetchTables["Fetch Table Schemas"]
    ValidatePluginSpec["Validate Plugin Specs"]
    
    ConfigFile --> ParseYAML
    ParseYAML --> ValidateConfig
    ValidateConfig --> InitPlugins
    InitPlugins --> FetchTables
    FetchTables --> ValidatePluginSpec
    
    subgraph "Validation Checks"
        CheckRequiredFields["Check Required Fields"]
        ValidateSourceConnections["Validate Source Connections"]
        ValidateDestConnections["Validate Destination Connections"]
        ValidateTableSelection["Validate Table Selection"]
        
        ValidateConfig --> CheckRequiredFields
        ValidatePluginSpec --> ValidateSourceConnections
        ValidatePluginSpec --> ValidateDestConnections
        ValidatePluginSpec --> ValidateTableSelection
    end
    
    ValidatePluginSpec --> |"Valid"| ReadyForSync["Ready for Sync"]
    ValidatePluginSpec --> |"Invalid"| ConfigError["Configuration Error"]
```

Sources: [cli/cmd/testdata/validate-config-error.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/validate-config-error.yml), [cli/cmd/testdata/test-connection-bad-connection.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/test-connection-bad-connection.yml)

## Common Configuration Issues

Here are some common configuration issues and their solutions:

| Issue | Description | Solution |
|-------|-------------|----------|
| Missing destination path | A destination is missing the required `path` field | Add the path field to the destination configuration |
| Invalid plugin version | Specified plugin version doesn't exist | Check available versions and update configuration |
| Invalid plugin specification | Plugin-specific configuration contains errors | Review plugin documentation for correct configuration |
| Plugin connection issues | Cannot connect to source or destination | Check connection parameters and credentials |
| Destination not found | Source references a destination that isn't defined | Ensure destination names match between source and destination blocks |

Sources: [cli/cmd/testdata/sync-missing-path-error.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/sync-missing-path-error.yml), [cli/cmd/testdata/validate-config-error.yml](https://github.com/cloudquery/cloudquery/blob/5064c039/cli/cmd/testdata/validate-config-error.yml)

## Environment Variable Substitution

CloudQuery supports environment variable substitution in configuration files using the `${VARIABLE_NAME}` syntax:

```yaml
kind: "destination"
spec:
  name: "test"
  path: "cloudquery/file"
  version: "v5.4.11"
  send_sync_summary: true
  spec:
    format: "json"
    path: ${CQ_FILE_DESTINATION}
```