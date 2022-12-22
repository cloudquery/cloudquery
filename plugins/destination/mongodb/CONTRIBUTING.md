# Contribution Guide to CloudQuery MongoDB Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

To run a MongoDB instance in a docker please run:

```bash
docker run -d -p 27017:27017 mongo
```

```bash
make test
```

## Lint

```bash
make lint
```