# Microsoft SQL Server destination plugin example

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";
import { Callout } from 'nextra-theme-docs'

<Badge text={"Latest: " + getLatestVersion("destination", "mssql")}/>

In this article we will show you a simple example of configuring Microsoft SQL Server destination plugin.

## Prerequisites

In order to be able to sync to Microsoft SQL Server you will need a running installation.
We will be using the [quickstart guide](https://learn.microsoft.com/en-us/sql/linux/quickstart-install-connect-docker)
for running Microsoft SQL Server locally using Docker.

### Create admin password for Microsoft SQL Server

Microsoft SQL Server enforces password complexity.
In order to successfully run the database you must specify a password that adheres the policy described
[here](https://learn.microsoft.com/en-us/sql/relational-databases/security/password-policy).

For this example we will be using `yourStrongP@ssword` as a password.

### Start Microsoft SQL Server locally

```sh
docker run \
  -e "ACCEPT_EULA=Y" \
  -e "MSSQL_SA_PASSWORD=yourStrongP@ssword" \
  -p 1433:1433 \
  -d \
  mcr.microsoft.com/mssql/server:2022-latest
```

### Create database to sync to

We will be using `cloudquery` for database name in this example.

```sh
docker exec $(docker ps -alq) \
  /opt/mssql-tools/bin/sqlcmd \
  -U "SA" \
  -P "yourStrongP@ssword" \
  -Q "CREATE DATABASE cloudquery;"
```

**Note**: `docker ps -alq` returns container ID for the latest started container.
You can use container ID discovered manually via `docker ps` output instead.

## Configure Microsoft SQL Server destination plugin

Once you've completed the steps from [Prerequisites](#prerequisites) section
you should be able to connect to the local `cloudquery` Microsoft SQL Server database
via the following connection string:

```text
server=localhost;user id=SA;password=yourStrongP@ssword;port=1433;database=cloudquery;
```

The (top level) spec section is described in the [Destination Spec Reference](/docs/reference/destination-spec).
The full configuration for Microsoft SQL Server destination plugin will look like this:

```yaml
kind: destination
spec:
  name:     "mssql"
  registry: "github"
  path:     "cloudquery/mssql"
  version:  "VERSION_DESTINATION_MSSQL"

  spec:
    connection_string: "server=localhost;user id=SA;password=yourStrongP@ssword;port=1433;database=cloudquery;"
```

<Callout type="info">
Make sure you use environment variable expansion in production instead of committing the credentials to the configuration file directly.
</Callout>
