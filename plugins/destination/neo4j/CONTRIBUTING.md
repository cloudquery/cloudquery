# Contribution Guide to CloudQuery Neo4j Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

To run the tests you will need a Neo4j database.

To run Neo4j docker on `neo4j://neo4j:test1234@localhost:7687` (see also official [docs](https://neo4j.com/developer/docker-run-neo4j/)):

```bash
docker run -p 7687:7687 -e NEO4J_AUTH=neo4j/test1234 -d neo4j:4.4
```

Once docker is up you can run:

```bash
make test
```

To change test database connection you can set the following environment variables:

- `CQ_DEST_NEO4J_CONNECTION_STRING`
- `CQ_DEST_NEO4J_USERNAME`
- `CQ_DEST_NEO4J_PASSWORD`

## Lint

```bash
make lint
```