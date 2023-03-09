# Contribution Guide to CloudQuery Neptune Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```


## Testing

To run the tests you will need a Gremlin database.

To run TinkerPop gremlin-server on `ws://localhost:8182` (see also official [docs](https://tinkerpop.apache.org/docs/current/reference/#gremlin-server-docker-image)):

```bash
docker run -p 8182:8182 -d tinkerpop/gremlin-server:3.6.2
```

Once docker is up you can run:

```bash
make test
```

To change test database connection you can set the following environment variables:

- `CQ_DEST_NEPTUNE_ENDPOINT` (default: `ws://localhost:8182`)
- `CQ_DEST_NEPTUNE_INSECURE` (if testing `wss://` with self-signed certificates)
- `CQ_DEST_NEPTUNE_USERNAME`
- `CQ_DEST_NEPTUNE_PASSWORD`

## Lint

```bash
make lint
```