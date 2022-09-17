# Overview

Sync command is the core command for CloudQuery. It does the following:

- Reads a given directory to read all configuration files
- Downloads all source and destination plugins specified in the configuration
- Start source plugins with the given config and sync the information extract by the source plugins to the specified destination plugins.

## Running

### Spawn or connection to PostgreSQL database

CloudQuery needs a PostgreSQL database (>=10). You can either spawn a local one (usually good for development and local testing)
or connect to an existing one.

For local, you can use the following docker command:

```bash
docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres
```

### Generate `cloudquery.yml`

An initial `cloudquery.yml` can be generated via `cloudquery init [provider]` (`provider` can be `aws`, `gcp` - see all options at [https://hub.cloudquery.io](https://hub.cloudquery.io)).

If you are using an existing database, you will have to update the `connection` section
in `cloudquery.yml`:

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
        database: postgres
        sslmode: disable
```

### Sync

Once `cloudquery.yml` is generated, run the following command to fetch the resources. (You need to be authenticated — see relevant section under each provider):

```powershell
cloudquery fetch
```
