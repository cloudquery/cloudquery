# Contribution Guide to CloudQuery Meilisearch Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

To run the tests, you will need a running Meilisearch instance.
You can use the [docker-compose.yaml](docker-compose.yaml) file to start one:

```bash
docker compose up -d
```

This will start a Meilisearch instance locally listening to port `7700`.

You will need to use an API key that has the following actions granted to run the plugin:

- `documents.add`
- `indexes.create`
- `indexes.get`
- `indexes.update`
- `tasks.get`
- `settings.get`
- `settings.update`
- `version`

However, the following additional actions are required to be granted to run tests:

- `search`

For testing purposes you can use the master key from [docker-compose.yaml](docker-compose.yaml)
(either the default `"test"` or the value from `MEILI_MASTER_KEY` env variable).

To run the tests:

```bash
make test
```

## Lint

```bash
make lint
```