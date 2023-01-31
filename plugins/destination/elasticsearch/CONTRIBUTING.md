# Contribution Guide to CloudQuery Elasticsearch Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

To run the tests, you will need a running Elasticsearch instance. You can use the docker-compose file in the root of the repository to start one, or use the `ELASTICSEARCH_ADDRESS` environment variable to point to an existing instance. The default is `localhost:9200`.

```bash
make test
```

## Lint

```bash
make lint
```