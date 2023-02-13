# Contribution Guide to CloudQuery PostgreSQL Source Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

To run the tests you will need a PostgreSQL database .

Running PostgreSQL docker `postgresql://postgresql:pass@localhost:5432/postgres?sslmode=disable`:

Note: You will need to add the `-c "wal_level=logical"` flag to enable logical replication, so the tests can run also the CDC scenario.

```bash
docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres:11 -c "wal_level=logical"
```

By default the tests will try to connect to `postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable`.
To change that you will need to set/export the environment variable `CQ_SOURCE_PG_TEST_CONN` prior to running the tests:

```bash
# for the above postgresql
# export CQ_SOURCE_PG_TEST_CONN="postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
make test
```

## Lint

```bash
make lint
```
