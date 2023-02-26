# Contribution Guide to CloudQuery MySQL Destination Plugin

## Running in Debug mode

Similar to all other CloudQuery plugins you can run the plugin in debug mode via:

```bash
go run main.go serve
```

## Setup

To run a test instance of MySQL you can use the following docker command:

```bash
docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=test -e MYSQL_DATABASE=cloudquery  -d mysql:5.7
```

> On Apple silicon you need to use the linux/amd64 platform via `--platform linux/amd64`

## Testing

```bash
make test
```

## Lint

```bash
make lint
```