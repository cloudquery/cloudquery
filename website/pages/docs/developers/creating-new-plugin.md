# Creating a New Plugin

This guide will help you write your own CloudQuery source plugin. CloudQuery's modular architecture means that a source plugin can be used to fetch data from any third-party API, and then be combined with a destination plugin to insert data into any supported destination. 

This guide assumes that you are somewhat familiar with CloudQuery. If you are not, we recommend starting by reading the [Quickstart guide](/docs/quickstart) and playing around with the CloudQuery CLI a bit first.

Though you by no means need to be an expert, you will also need some familiarity with Go. The official [Go Tutorial](https://go.dev/doc/tutorial/getting-started) and [A Tour of Go](https://go.dev/tour/welcome/1) are great resources to learn the basics and prepare your environment.

## Core Concepts

Before we dive in, let's quickly cover some core concepts of CloudQuery plugins, so that they're familiar when we see our first example.

### Syncs

A sync is the process that gets kicked off when a user runs `cloudquery sync`. A sync is responsible for fetching data from a third-party API and inserting it into the destination (database, data lake, stream, etc.). When you write a source plugin for CloudQuery, you will only need to implement the part that interfaces with the third-party API. The rest of the sync process, such as delivering to the destination database, is handled by the CloudQuery SDK.

### Tables and Services

A **table** is the term CloudQuery uses for a collection of related data. In most databases it directly maps to an actual database table, but in some destinations it could be stored as a file, stream or other medium. Inside plugin code, tables get grouped into collections called "services". Many REST APIs are logically grouped, and services are meant to map closely to these underlying API groupings. For example, an API might expose an endpoint called `GET /v1/accounts/users`. The service in this case would be called `accounts`, and the table `users`. The final table name will be `<plugin_name>_<service_name>_<table_name>`, e.g. `myplugin_accounts_users`. 

Services each get their own directory under the `services` directory of your plugin. Inside a service directory, every table will typically have its own `.go` file. A table is defined as a function that returns a [`*schema.Table`](https://github.com/cloudquery/plugin-sdk/blob/a1409ac07858d9b1dca5098e430469c943bf8b63/schema/table.go#L34-L70).

Not every plugin will have enough tables to justify grouping them into services. For plugins with only a few tables, it's fine to put them directly in the `resources` directory. We will look at examples of this soon! For now, let's cover a few more important concepts.

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

Incremental tables are great for efficiency, but add some additional complexity both on you and on your users. As the plugin author, you should consider first whether the table needs to be incremental, then whether it can be made to be incremental. You can also [read more about managing incremental tables](/docs/advanced-topics/managing-incremental-tables).

## Creating Your First Plugin

In this section we will go through all the steps of building a simple source plugin. We will start by creating a new plugin from scratch, then we will add a table to it. To serve as a fun real-world example, we will create a plugin that fetches comic data from the [XKCD API](https://xkcd.com/json.html).

### Initializing Your Plugin with the `scaffold` Tool

The easiest way to get started writing a plugin is to use the `scaffold` tool. This tool will create a new plugin directory with all the boilerplate code you need to get started. It will also create a `services` directory with an example table.

The scaffold tool is available as a binary for Linux, macOS and Windows. You can download the latest version from the [releases page](https://github.com/cloudquery/cloudquery/releases?q=scaffold&expanded=true).

On MacOS, you can install the tool using Homebrew:

```bash
brew install cloudquery/tap/scaffold
```

With the tool installed, you can create a new plugin by running (replace `<org>` and `<name>` with values for your GitHub org and the name of your plugin):

```bash
cq-scaffold source <org> <name>
```

This will create a new directory called `cq-source-<name>`. You should then `cd` into the directory and run `go mod tidy` to download the dependencies.

At the time of writing, the scaffold creates a directory structure that looks like this:

```text
.
├── README.md
├── client
│   ├── client.go
│   └── spec.go
├── cq-source-xkcd
├── go.mod
├── go.sum
├── main.go
├── plugin
│   └── plugin.go
└── resources
    └── table.go
```

### Creating a Table

The scaffold tool creates a single table in the `resources` directory. Let's take a look at the code in `resources/table.go` that was generated for a new XKCD source plugin:

```go
package resources

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SampleTable() *schema.Table {
	return &schema.Table{
		Name:     "xkcd_sample_table",
		Resolver: fetchSampleTable,
		Columns: []schema.Column{
			{
				Name: "column",
				Type: schema.TypeString,
			},
		},
	}
}

func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return fmt.Errorf("not implemented")
}
```

In this example, we have a table called `xkcd_sample_table` with a single column called `column`. The `Resolver` field contains the resolver function that will be called to populate the table with data. The `fetchSampleTable` function is a placeholder that returns an error. Our job as plugin authors will be to add the correct columns for the table, and implement the resolver function.

### Adding Columns to the Table

Adding columns to a table is easy, as long as you have a Go struct. The CloudQuery plugin SDK will automatically map the fields of the struct to the columns of the table. In many cases an existing Go SDK will provide you with this struct. Then we can add a `Transform` property that calls `transformers.TransformWithStruct(&<StructName>{})` with a pointer to the struct. This will automatically map the fields of the struct to the columns of the table. For our hypothetical XKCD plugin, we don't have an SDK to work with, so we will create our own struct inside a new `xkcd` package. The struct will look like this:

```go
type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}
```

We'll also rename the `SampleTable` function to `Comics`, update some properties and add the `Transformer` property: 

```go
func Comics() *schema.Table {
	return &schema.Table{
		Name:     "xkcd_comics",
		Resolver: fetchComics,
		Transform: transformers.TransformWithStruct(&xkcd.Comic{}),
	}
}
```

### Writing a Table Resolver

With the columns defined, we can now write the resolver function. The resolver function is responsible for fetching the data from the API and returning it to CloudQuery. The resolver function takes a `context.Context` object, a `schema.ClientMeta` object, a `*schema.Resource` object, and a `chan<- interface{}` object. The `context.Context` object is used to cancel the resolver function if the user cancels the sync. The `schema.ClientMeta` object is a generic object that can be used to store any data that needs to be shared between resolvers. The `*schema.Resource` object is the parent resource of the table, if any, and is used to implement parent-child relationships. In our case this will be `nil`. The `chan<- interface{}` object is used to send the data back to CloudQuery.

We won't go into all the details of making API calls here, but let's look at what a resolver function for our XKCD plugin might look like:

```go

func fetchComics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*Client)
	latest, err := client.XKCD.GetLatestComic(ctx)
	if err != nil {
		return err
	}
	res <- latest

	for i := 1; i < latest.Num; i++ {
		comic, err := client.XKCD.GetComic(ctx, i)
		if err != nil {
			return err
		}
		res <- comic
	}
	return nil
}
```

In the above code, we are getting a list of Comics from the XKCD API and sending them to the CloudQuery over the `res` channel. We first need to get the latest comic, then we can iterate through all the IDs from 1 to that number. You can send items to the channel one at a time, or as a slice of items. The sooner an is dispatched over the channel, the sooner it will be written to the destination(s), so we prefer to write them as soon as they are available. And as long as the struct sent matches the one used for the table, the CloudQuery SDK will handle the rest.

In the above example, we used a `Client` struct that we haven't talked about yet. The `Client` struct is used to store any data that needs to be shared between resolvers. For example, it may store the API key that we'll need to make API calls, or an SDK client that we'll use to make API calls. The `Client` struct is defined in the `client` directory, and is instantiated with a call to `client.New` in the `plugin` directory. In this case, we were using it to store an instance of the XKCD client. (We won't show the full XKCD client implementation here.)
 
### Testing the Plugin

There are two options for running a plugin before as a developer before it is released: as a gRPC server, or as a standalone binary. We will briefly summarize both options here, or you can read about them in more detail in [Running Locally](/docs/developers/running-locally).

#### Run the Plugin as a gRPC Server

This mode is especially useful for setting breakpoints your code for debugging, as you can run it in server mode from your IDE and attach a debugger to it. To run the plugin as a gRPC server, you can run the following command in the root of the plugin directory:

```shell
go run main.go serve
```

(Note: If you see errors about missing dependencies, you can run `go mod tidy` to fix them.)

This will start a gRPC server on port 7777. You can then create a config file that sets the `registry` and `path` properties to point to this server. For example:

```yaml copy filename=config.yaml
kind: source
spec:
  name: "xkcd"
  registry: "grpc"
  path: "localhost:7777"
  version: "v1.0.0"
  tables: 
    ["*"]
  destinations:
    - "sqlite"
---
kind: destination
spec:
  name: sqlite
  path: cloudquery/sqlite
  version: "v1.2.1"
  spec:
    connection_string: ./db.sql
```

With the above configuration, we can now run `cloudquery sync` as normal:

```shell copy
cloudquery sync config.yaml
```

Note that when running a source plugin as a gRPC server, errors with the source plugin will be printed to the console running the gRPC server, *not* to the CloudQuery log like usual. 

#### Run the Plugin as a Standalone Binary

To run the plugin as a standalone binary, you can run the following command in the root of the plugin directory:

```shell copy
go build
```

This will create a binary with the name of the plugin directory (so, `cq-source-<plugin-name>`). We can then refer to this binary by setting the registry to `local` and `path` as the path to the binary. Example:

```yaml copy filename=config.yaml
kind: source
spec:
  name: "xkcd"
  registry: "local"
  path: "/path/to/cq-source-xkcd"
  version: "v1.0.0"
  tables: 
    ["*"]
  destinations:
    - "sqlite"
---
kind: destination
spec:
  name: sqlite
  path: cloudquery/sqlite
  version: "v1.2.1"
  spec:
    connection_string: ./db.sql
```

With the above configuration, we can now run `cloudquery sync` as normal:

```shell copy
cloudquery sync config.yaml
```

This time errors will be logged to `cloudquery.log`, as usual. This mode is closest to how the plugin will run when it is released, as the CLI is in charge of managing the plugin process.

### Writing a Column Resolver

Sometimes it is necessary, or useful, to add some additional information to a table. This doesn't happen often, however, and for the XKCD plugin we will need to come up with a contrived example to show how this works. Let's imagine that, in addition to the `Comic` struct fields, we also want to add whether the comic is a "good" comic or not. We can do this by adding a new column to the table, and then writing a resolver function for that column. The column will be called `is_good` and will be a boolean. We'll add the column to the table definition like this:

```go
func Comics() *schema.Table {
	return &schema.Table{
		Name:     "xkcd_comics",
		Resolver: fetchComics,
		Transform: transformers.TransformWithStruct(&xkcd.Comic{}),
		Columns: []schema.Column{
			{
				Name:     "is_good",
				Type:     schema.TypeBool,
				Resolver: resolveComicIsGood,
			},
		},
	}
}
```

The `Resolver` property is the function that will be called to resolve the column value. We'll define that function next:

```go
func resolveComicIsGood(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	comic := resource.Item.(xkcd.Comic)
	resource.Set(c.Name, strings.Contains(comic.Title, "XKCD"))
	return nil
}
```

As big fans of meta-jokes, we define only comics with "XKCD" in the title to be good. 

### Adding Multiplexing

For our simple XKCD plugin, multiplexing is not necessary. But let's say we were writing a plugin that can fetch from multiple accounts. In that case, we may define an `AccountMultiplex` multiplexer inside a new `multiplexers.go` file in the `client` directory: 

```go
func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, acc := range client.accounts {
		l = append(l, client.WithAccount(acc))
	}
	return l
}
```

This also requires a new `WithAccount` method on the Client struct that sets an Account property on the client:

```go
func (c *Client) WithAccount(account string) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("account", account).Logger()
	newC.Account = account
	return &newC
}
```

It is also important to update the `ID()` method on the client to include the account name. This is used in logging and error messages to identify the client, but also internally in the SDK to identify the client. We can update the `ID()` method to include the account name like this

```go
func (c *Client) ID() string {
	return fmt.Sprintf("myplugin:%s", c.Account)
}
```

The exact format doesn't matter, as long as it is unique for every multiplexed value. Some plugins also include `spec.Name` in the ID, to help identify the plugin in scenarios where multiple instances are run in parallel.

Now we can instruct the plugin SDK to use this multiplexer, where appropriate, by setting the `Multiplex` property on the table to `client.AccountMultiplex`:

```go
func MyTable() *schema.Table {
	return &schema.Table{
		Name:     "sample_table",
		Resolver: fetchSampleTable,
		Multiplex: client.AccountMultiplex,
		// other properties ...
	}
}
```

Inside the `fetchSampleTable` resolver, we would then be able to get the current Account by accessing the `Account` property on the client:

```go
func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*Client)
	account := client.Account
	// ...
}
``` 

The [GitHub plugin multiplexers](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/github/client/multiplexer.go) can serve as a good example of how to implement and use multiplexing. In that case, some tables multiplex on organization, while others multiplex on organization and repository combined.

<!-- TODO
## Testing Strategies
-->

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

## Resources

- [Discord](https://cloudquery.io/discord)
- [How to Write a CloudQuery Source Plugin](https://www.youtube.com/watch?v=3Ka_Ob8E6P8) (Video 🎥)

