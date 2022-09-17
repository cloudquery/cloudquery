import { Callout } from 'nextra-theme-docs'

# Overview

This is an overview of the configuration options of CloudQuery's main configuration file `cloudquery.yml`.

## Main cloudquery block

The `cloudquery` block must be specified exactly once per `cloudquery.yml`. This usually looks like:

```yaml
cloudquery:
  providers:
    - name: aws
      version: latest
  connection:
    type: postgres
    username: postgres
    password: pass
    host: localhost
    port: 5432
    schema: public
    database: postgres
    sslmode: disable
```

### `cloudquery.connection`

A map of values that defines the connections details to your PostgreSQL database.

- **`type`** **(required)** - Type of database that CloudQuery will connect to. Only valid value is `postgres`
- **`username`** **(required)** - Username that CloudQuery will use when interacting with the Postgres database
- **`password`** **(required)** - Password for user that CloudQuery will use to authenticate into the Postgres database
- **`host`** **(required)** - Hostname or IP address of Postgres database
- **`port`** **(optional)** - Port of the Postgres database that CloudQuery will connect to. Default value is `5432`
- **`schema`** **(optional)** - The name of the schema that CloudQuery will use. Default value is `public`
- **`database`** **(required)** - Name of the Postgres database that CloudQuery will connect to
- **`sslmode`** **(required)** - Postgres setting for specifying the level of security you want to enforce in the connection between CloudQuery and your database. If you are running CloudQuery locally in a docker container the typical value is `disable`. Other valid options include: `allow`, `prefer`, `require`, `verify-ca`, `verify-full`

### `cloudquery.providers`

A list of objects that defines which providers and the corresponding versions that CloudQuery should download and ensure are downloaded and ready to be invoked:

- **`name`** - Name of the provider you want to use. Should be in the form `organization/name` if no organization is set then it will assume the organization is `cloudquery`
- **`source`** **(Optional)** - By default CloudQuery will assume the location is `github.com/organization/cq-provider-<name>` (where the `<name>` comes from the `name` attribute) unless user specifies a different location.
- **`version`** - Based on Git tags of the repository. User can define either a specific tagged version or `latest`

## Main Providers Block

The `providers` block at the root of the file must be defined exactly once. It specifies all of the provider specific configurations.

```yaml
providers:
  - name: <provider-name>
    configuration:
      \\ This will be provider specific configurations
    alias: <unique_identifier>
    resources:
      - "*"
    skip_resources:
      - "slow.resource_1"
  - name: <provider-name>
    configuration:
      \\ This will be provider specific configurations
    alias: <unique_identifier_2>
    resources:
      - "*"
```

### `providers`

A list of objects that represent a provider that will be configured.

Each provider has the following blocks that can be set:

- `name` - The name of the provider that corresponds to a named provider specified in `cloudquery.providers`
- `configuration` - The arguments are different from provider to provider and their documentation can be found in [CloudQuery Hub](https://hub.cloudquery.io).
- `resources` - A list of resources to fetch configuration and metadata for. You can specify all supported resources by providing `*` as the first value.
- `alias` **(Optional)** - A unique identifier for the provider so that you can have multiple instances for the same provider
- `max_goroutines` **(Optional)** - The maximum number Go routines created by cloudquery for the purpose of parallel resource fetching. This is useful for providers that have low concurrency thresholds and for compute resources that have minimal memory and CPU available. Value must be an integer greater than 0. Default behavior is that CloudQuery will attempt to use all resources available but will try and factor in CPU and File descriptor limitations.
- `max_parallel_resource_fetch_limit` **(Optional)** - The maximum number of resources that are attempted to be fetched in parallel. This is useful for providers that have low concurrency thresholds and for compute resources that have minimal memory and CPU available Value must be an integer greater than 0. Default behavior is CloudQuery will attempt to fetch as many resources in parallel as possible.
- `resource_timeout` **(Optional)** - The number of seconds that CloudQuery will spend fetching any single resource. Value must be an integer greater than 0. Default behavior is unlimited timeout.
- `skip_resources` **(Optional)** - A list of resources that should be explicitly skipped. This can help if you are using a `*` or dynamically generating the `resources` value.

<Callout type="info">

You can have multiple providers of the same type specified here as long as you specify an alias like the example above. This can be useful if you want to fetch data with different rate limiting parameters.

</Callout>
