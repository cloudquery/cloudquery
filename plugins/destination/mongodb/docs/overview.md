---
name: MongoDB
stage: GA
title: MongoDB Destination Plugin
description: CloudQuery MongoDB destination plugin documentation
---
# MongoDB Destination Plugin

:badge

This destination plugin lets you sync data from any CloudQuery source to a MongoDB database.

Supported database versions:

- MongoDB >= 3.6 (The same minimum version supported by the official [Go driver](https://github.com/mongodb/mongo-go-driver))

## Configuration

### Example

:configuration

:::callout{type="info"}
Make sure to use [environment variable substitution](/docs/advanced-topics/environment-variable-substitution) in production instead of committing the credentials to the configuration file directly.
:::

The MongoDB destination utilizes batching, and supports [`batch_size`](/docs/reference/destination-spec#batch_size) and [`batch_size_bytes`](/docs/reference/destination-spec#batch_size_bytes). 

### MongoDB Spec

This is the (nested) spec used by the MongoDB destination Plugin.

- `connection_string` (`string`) (required)

  MongoDB URI as described in the official MongoDB [documentation](https://www.mongodb.com/docs/manual/reference/connection-string/).
  
  Example connection strings:
  
  - `"mongodb://username:password@hostname:port/database"` basic connection
  - `"mongodb+srv://username:password@cluster.example.com/database"` connecting to a MongoDB Atlas cluster
  - `"mongodb://localhost:27017/myDatabase?authSource=admin"` specify authentication source

- `database` (`string`) (required)

  Database to sync the data to.

- `batch_size` (`integer`) (optional) (default: `1000`)

  Maximum amount of items that may be grouped together to be written in a single write.

- `batch_size_bytes` (`integer`) (optional) (default: `4194304` (= 4 MiB))

  Maximum size of items that may be grouped together to be written in a single write.
