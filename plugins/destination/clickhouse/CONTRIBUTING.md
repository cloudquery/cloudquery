# Contribution Guide to CloudQuery ClickHouse Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

To run the tests, you will need a running ClickHouse instance.
You can use the [docker-compose.yaml](docker-compose.yaml) file to start one:

```bash
docker-compose up -d
```

This will start a ClickHouse instance and create the `cloudquery` database along with user `cq` (password: `test`).

To run the tests:

```bash
make test
```

## Lint

```bash
make lint
```