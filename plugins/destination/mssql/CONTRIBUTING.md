# Contribution Guide to CloudQuery Microsoft SQL Server Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

Tests require an SQL Server instance available.

You can run the following script to start the local server & create the required database:

```bash
docker run \
  --platform=linux/amd64 \
  -e ACCEPT_EULA=Y \
  -e MSSQL_PID=Express \
  -e MSSQL_SA_PASSWORD='yourStrongP@ssword' \
  -e DB_USER=SA \
  -e DB_NAME=cloudquery \
  -p 1433:1433 \
  -d mcr.microsoft.com/mssql/server:2017-latest
```

After that, ensure that the database is created:

```bash
docker exec $(docker ps -alq) \
  /opt/mssql-tools/bin/sqlcmd \
  -U "SA" \
  -P 'yourStrongP@ssword' \
  -Q "CREATE DATABASE cloudquery;"
```

Then you can run tests:

```bash
make test
```

## Lint

```bash
make lint
```