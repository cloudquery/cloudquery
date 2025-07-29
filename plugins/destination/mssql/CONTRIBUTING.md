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
docker compose up --wait
```

Then you can run tests:

```bash
make test
```

## Lint

```bash
make lint
```

To clean up:

```bash
docker compose down --volumes
```