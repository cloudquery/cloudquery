# Contribution Guide to CloudQuery Elasticsearch Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

To run the tests, you will need a running Elasticsearch instance. You can use the docker-compose file in the root of the repository to start one:
 
```bash
docker compose up -d
```

This will also start a Kibana instance at `localhost:5601`. Or use the `ELASTICSEARCH_ADDRESS` environment variable to point to an existing instance. The default is `localhost:9200`.

To run the tests:

```bash
make test
```

## Lint

```bash
make lint
```
