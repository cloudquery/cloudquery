# Contribution Guide to CloudQuery OracleDB Source Plugin

## Running in Debug mode

Similar to all other CloudQuery plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Setup

To run a test instance of OracleDB you can use the following docker command:

```bash
docker run --name oracledb -p 1521:1521 -e ORACLE_DATABASE=cloudquery -e APP_USER=cq -e APP_USER_PASSWORD=test -e ORACLE_PASSWORD=test gvenzl/oracle-xe
```

> For Apple silicon see [here](https://github.com/gvenzl/oci-oracle-xe/tree/ea0533d5f6778f6385941fec239a71af75605976#oracle-xe-on-apple-m-chips)

## Testing

```bash
make test
```

## Lint

```bash
make lint
```

