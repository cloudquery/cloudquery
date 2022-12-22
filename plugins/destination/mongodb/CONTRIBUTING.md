# Contribution Guide to CloudQuery MongoDB Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

To run a mongodb instance in a docker please run:

```bash
run -d -p 27017:27017 mongo
```

```bash
make test
```

## Lint

```bash
make lint
```