# Contribution Guide to CloudQuery BigQuery Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

```bash
make test
```

## Lint

```bash
make lint
```

## Running with a local emulator

Run the [emulator](https://github.com/goccy/bigquery-emulator): 

```bash
docker run -d -p 9050:9050 -it ghcr.io/goccy/bigquery-emulator:latest --project=test --dataset=test
```

In the spec file, set `project_id` and `dataset_id` to be `test` and set `endpoint` to be `http://localhost:9050`:

```yaml copy
spec:
    kind: destination
    # other opts...
    spec:
        project_id: test
        dataset_id: test
        endpoint: http://localhost:9050
```
