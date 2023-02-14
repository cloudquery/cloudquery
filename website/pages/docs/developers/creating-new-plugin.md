# Creating a New Plugin

This guide will help you write your own CloudQuery source plugin. CloudQuery's modular architecture means that a source plugin can be used to fetch data from any third-party API, and then be combined with a destination plugin to insert data into any supported destination. 

This guide assumes that you are somewhat familiar with CloudQuery. If you are not, we recommend starting by reading the [Quickstart guide](/docs/quickstart) and playing around with the CloudQuery CLI a bit first.

Though you by no means need to be an expert, you will also need some familiarity with Go. The official [Go Tutorial](https://go.dev/doc/tutorial/getting-started) and [A Tour of Go](https://go.dev/tour/welcome/1) are great resources to learn the basics and prepare your environment.

## Core Concepts

This section will cover some core concepts of CloudQuery plugins.

### Syncs

A sync is the process that gets kicked off when a user runs `cloudquery sync`. A sync is responsible for fetching data from a third-party API and inserting it into the destination (database, data lake, stream, etc.). When you write a source plugin for CloudQuery, you will only need to implement the part that interfaces with the third-party API. The rest of the sync process, such as delivering to the destination database, is handled by the CloudQuery SDK.

### Tables and Services

A **table** is the term CloudQuery uses for a collection of related data. In most databases it directly maps to an actual database table, but in some destinations it could be stored as a file, stream or other medium. Inside plugin code, tables get grouped into collections called "services". Many REST APIs are logically grouped, and services are meant to map closely to these underlying API groupings. For example, an API might expose an endpoint called `GET /v1/accounts/users`. The service in this case would be called `accounts`, and the table `users`. The final table name will be `<plugin_name>_<service_name>_<table_name>`, e.g. `myplugin_accounts_users`. 

Services each get their own directory under the `services` directory of your plugin. Inside a service directory, every table will typically have its own `.go` file. A table is defined as a function that returns a [`*schema.Table`](https://github.com/cloudquery/plugin-sdk/blob/a1409ac07858d9b1dca5098e430469c943bf8b63/schema/table.go#L34-L70). We will look at examples of this soon! For now, let's cover a few more important concepts.

### Resolvers

Resolvers are functions associated with a table that get called when it's time to populate data for that table. There are two types of resolvers:

#### Table resolvers

Table resolvers are responsible for fetching data from the third-party API. A table resolver receives a `res` (results) channel as argument, to which it should send all results from the API. For top-level tables, this function will only be called once per multiplexer client. For dependent tables, the resolver will be called once for each parent row, and the parent resource will be passed in as well. (More on this, and multiplexers, shortly.)

#### Column resolvers

Column resolvers are responsible for mapping data from the third-party API into the columns of the table. In most cases, you will not need to implement this, as the SDK will automatically map data from the struct passed in by the table resolver to the columns of the table. But in some cases, you may need to implement a custom column resolver to fetch additional data or do custom transformations.

### Multiplexers

Multiplexers are a way to parallelize the fetching of data from the third-party API. Some top-level tables require multiple calls to fetch all their data. For example, a sync for the GitHub source plugin that fetches data for multiple organizations, will need to make one call per organization to list all repositories. By multiplexing over organizations, these top-level queries can also be done in parallel. Each table defines the multiplexer that it should use. The CloudQuery plugin SDK will then call the table resolver once for each client in the multiplexer.

### Incremental Tables

Some APIs lend themselves to being synced incrementally: rather than fetch all past data on every sync, an incremental table will only fetch data that has changed since the last sync. This is done by storing some metadata in a state **backend**. The metadata is known as a **cursor**, and it marks where the last sync ended, so that the next sync can resume from the same point. Incremental syncs can be vastly more efficient than full syncs, especially for tables with large amounts of data. This is because only the data that's changed since the last sync needs to be retrieved, and in many cases this is a small subset of the overall dataset.

Incremental tables are great for efficiency, but add some additional complexity both on you and on your users. As the plugin author, you should consider first whether the table needs to be incremental, then whether it can be made to be incremental.

## Creating Your First Plugin

### Initializing Your Plugin with the `scaffold` Tool

The easiest way to get started writing a plugin is to use the `scaffold` tool. This tool will create a new plugin directory with all the boilerplate code you need to get started. It will also create a `services` directory with an example table.

The scaffold tool is available as a binary for Linux, macOS and Windows. You can download the latest version from the [releases page](https://github.com/cloudquery/cloudquery/releases?q=scaffold&expanded=true).

On MacOS, you can install the tool using Homebrew:

```bash
brew install cloudquery/tap/scaffold
```

With the tool installed, you can create a new plugin by running (replace `<org>` and `<name>` with values for your GitHub org and the name of your plugin):

```bash
scaffold source <org> <name>
```

This will create a new directory called `cq-source-<name>`. You should then `cd` into the directory and run `go mod tidy` to download the dependencies.

At the time of writing, the scaffold creates a directory structure that looks like this:

```text
.
├── README.md
├── client
│   ├── client.go
│   └── spec.go
├── cq-source-tes
├── go.mod
├── go.sum
├── main.go
├── plugin
│   └── plugin.go
└── resources
    └── table.go
```

### Creating a Table

The scaffold tool creates a single table in the `resources` directory. Let's take a look at the code:

```go
```

<TODO>

### Writing a Table Resolver

<TODO>

### Writing a Column Resolver

<TODO>

### Adding Multiplexing

<TODO>

## Testing Strategies

<TODO>

## Releasing and Deploying Your Plugin

Releasing a plugin for use by the wider CloudQuery community involves two steps:

1. Create a public GitHub repository for your plugin named `cq-source-<plugin-name>` and push your code to it. 
2. Create a release in GitHub. The release name should be the version of your plugin, e.g. `v0.1.0`. More information about managing releases can be found in [GitHub's documentation](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository).

Users can then import your plugin by specifying the GitHub repository path in their config file, but leaving out the `cq-source-` part, e.g.:

```yaml copy
kind: source
spec:
  name: cloudwidgets
  path: myorg/<plugin-name>
```

This will import the plugin from the repository hosted at `github.com/myorg/cq-source-<plugin-name>`.

### Naming Conventions

Community plugins use the following GitHub repository naming conventions:

- `org/cq-source-<name>` for source plugins
- `org/cq-destination-<name>` for destination plugins

A community plugin using this convention can be imported in a config by using:

```yaml copy
kind: source
spec:
  path: org/name
```

for source plugins, or

```yaml copy
kind: destination
spec:
  path: org/name
```

for destination plugins.

Names should not contain dashes or underscores. So for example, if you are developing a source plugin for a new cloud service called Cloud Widgets, you should create the plugin repository under `org/cq-source-cloudwidgets`.

Official plugins, in contrast, are contained in the [CloudQuery repository](https://github.com/cloudquery/cloudquery/tree/main/plugins). By convention, they can be imported using a special path `cloudquery/<name>`, e.g.:

```yaml copy
kind: source
spec:
  path: cloudquery/aws
```

## Real-world Examples

A good way to learn how to create a new plugins is to look at the following examples:

- The [GitHub Source Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/github) is an example of a medium-complexity plugin with concepts you are likely to already be familiar with.
- The [Hacker News Source Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/hackernews) is a good example of a plugin with incremental tables.
- The [GCP Source Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/gcp) is a good example of a complex plugin with code generation elements and test server-based testing.

This guide doesn't cover destination plugins yet , but you can also look at the following examples:
- The [PostgreSQL Destination Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/destination/postgresql) is a good example of an "unmanaged" destination that handles batching itself
- The [BigQuery Destination Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/destination/bigquery) is a good example of a "managed" destination that writes to each table in separate batches

Other source and destination plugins to reference can be found [here](https://github.com/cloudquery/cloudquery/tree/main/plugins)
