# CloudQuery Plugin Configuration

A CloudQuery sync fetches data from cloud accounts (sources) and writes it to one or more destinations. 
A sync requires at least one source- and one destination configuration. 
Configuration files are specified in YAML format and can be either split across multiple files or combined.

## Example using multiple files

One option is to maintain configuration for your source and destination plugins in separate files.

Here is a simple example with only one source and one destination plugin:

```yaml title="aws.yml"
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  tables: ["*"]
  destinations: ["postgresql"]
```

```yaml title="postgresql.yml"
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "VERSION_DESTINATION_POSTGRESQL"
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```

With these two files, we can run a sync using:

```shell copy
cloudquery sync aws.yml postgresql.yml
```  

### Adding another source

Let's imagine we now wanted to add a `gcp` source as well. We can add its configuration in a new file:

```yaml title="gcp.yml"
kind: source
spec:
  name: gcp
  path: cloudquery/gcp
  version: "VERSION_SOURCE_GCP"
  tables: ["*"]
  destinations: ["postgresql"]
```

And now sync both `aws` and `gcp` to `postgresql` in a single command: 

```shell copy
cloudquery sync aws.yml gcp.yml postgresql.yml
``` 

## Example using one file

You can also combine sources and destinations into a single file by separating the sections with `---`:

```yaml title="config.yml"
kind: source
spec:
  name: aws
  path: cloudquery/aws
  version: "VERSION_SOURCE_AWS"
  tables: ["*"]
  destinations: ["postgresql"]
---
kind: destination
spec:
  name: postgresql
  path: cloudquery/postgresql
  version: "VERSION_DESTINATION_POSTGRESQL"
  spec:
    connection_string: ${PG_CONNECTION_STRING}
```

Now we can run a sync using:

```shell copy
cloudquery sync config.yml
```

This example shows only two plugin sections, but a config file is allowed to contain any number of plugin sections. 