# Contribution Guide to CloudQuery Snowflake Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

```bash
# export SNOW_TEST_DSN="username:password@account_locator.europe-west4.gcp/testdb/public?warehouse=test
make test
```


## Lint

```bash
make lint
```