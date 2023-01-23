# Contribution Guide to CloudQuery Kafka Destination Plugin

## Running in Debug mode

Similar to all other CQ plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Testing

To run the tests you will need a Kafka cluster

Running a local Kafka cluster

```bash
docker-compose up
```

```bash
# If your kafka cluster you are testing is located elsewhere you can change the following
# export CQ_DEST_KAFKA_CONNECTION_STRING="localhost:29092"
# This is only needed if SASL is used
# export CQ_DEST_KAFKA_SASL_USERNAME=""
# export CQ_DEST_KAFKA_SASL_PASSWORD=""
make test
```

## Lint

```bash
make lint
```
