# Contribution Guide to CloudQuery MongoDB Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

The plugin can be tested against either a standalone instance or a single-node replica set. CI runs both topologies.

### Standalone

```bash
docker compose -f docker-compose.standalone.yaml up -d --wait
export CQ_DEST_MONGODB_TEST_CONN="mongodb://localhost:27017"
make test
```

### Single-node replica set

The replica set compose file initiates `cloudquery` automatically via an idempotent healthcheck.

```bash
docker compose -f docker-compose.replicaset.yaml up -d --wait
export CQ_DEST_MONGODB_TEST_CONN="mongodb://localhost:27017/?replicaSet=cloudquery"
make test
```

To tear down: `docker compose -f <file> down -v`.

## Lint

```bash
make lint
```