# Contribution Guide to CloudQuery PostgreSQL Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

To run the tests you will need a PostgreSQL database (and/or a CockroachDB - both are tested in CI).

Running PostgreSQL docker `postgresql://postgresql:pass@localhost:5432/postgres?sslmode=disable`:

```bash
docker run -p 5432:5432 -e POSTGRES_PASSWORD=pass -d postgres:11
```

Running CockroachDB docker on `postgresql://root@localhost:26257/postgres?sslmode=disable`:

```bash
docker run -p 26257:26257  cockroachdb/cockroach:v22.1.8  start-single-node  --insecure
```

By default the tests will try to connect to `postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable`.
To change that you will need to set/export the environment variable `CQ_DEST_PG_TEST_CONN` prior to running the tests:

```bash
# for the above postgresql
# export CQ_DEST_PG_TEST_CONN="postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
# for the above cockroachdb
# export CQ_DEST_PG_TEST_CONN="postgresql://root@localhost:26257/postgres?sslmode=disable"
make test
```

## Lint

```bash
make lint
```
