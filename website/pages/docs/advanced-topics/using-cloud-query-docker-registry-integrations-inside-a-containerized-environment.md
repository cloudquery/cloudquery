# Using CloudQuery Docker Registry Integrations Inside a Containerized Environment 

CloudQuery CLI uses the [Docker CLI](https://docs.docker.com/engine/reference/commandline/cli/) and [Engine API](https://docs.docker.com/engine/api/) to run Docker integrations.
When using the CloudQuery CLI Docker image, Docker integrations don't work out of the box, as the Docker CLI and Engine API are not available in the container by default.
This guide will show you how to run Docker integrations with the CloudQuery CLI Docker image using Docker Compose.

## Prerequisites

- [Docker installed (with Docker Compose)](https://docs.docker.com/get-docker/)
- [CloudQuery CLI](https://docs.cloudquery.io/docs/quickstart)

## Setup

1. Run `cloudquery login` to authenticate with the CloudQuery registry.
2. Create a file named `spec.yml` with the configuration for the Docker integration. We will use this spec file to pull the Docker integration image locally from the private CloudQuery Docker registry. The `airtable` and `postgresql` integrations are used as an example.

```yaml filename="spec.yml"
kind: source
spec:
  name: airtable
  path: cloudquery/airtable
  version: VERSION_SOURCE_AIRTABLE
  tables: ["*"]
  destinations: ["postgresql"]
---
kind: destination
spec:
  name: "postgresql"
  path: "cloudquery/postgresql"
  version: "VERSION_DESTINATION_POSTGRESQL"
```
3. Run `cloudquery plugin install spec.yml` to pull the Docker integration image locally.

## Running a Sync

1. [Create a CloudQuery API Key](https://docs.cloudquery.io/docs/deployment/generate-api-key) to be used with the Docker Compose file.
2. Create a `docker-compose.yml` file with the following content. The file configures the CLI docker image, the Docker integration image, a PostgreSQL database and a configuration spec that sets up a connection between the CLI and the Docker integration using gRPC.

```yaml filename="docker-compose.yml"
version: '3.1'
services:
  cli:
    container_name: cli
    image: ghcr.io/cloudquery/cloudquery:latest
    command: ["sync", "/spec.yml", "--log-console", "--log-format", "json"]
    environment:
      CLOUDQUERY_API_KEY: ${CLOUDQUERY_API_KEY}
      # We reference this environment variable in the `spec.yml` config block below
      # Other plugins will require other environment variables
      AIRTABLE_ACCESS_TOKEN: ${AIRTABLE_ACCESS_TOKEN}
    configs:
      - spec.yml
    depends_on:
      airtable:
        condition: service_healthy
      postgres:
        condition: service_healthy
  airtable:
    container_name: airtable
    image: docker.cloudquery.io/cloudquery/source-airtable:VERSION_SOURCE_AIRTABLE
    # We use `cloudquery login` and `cloudquery plugin install spec.yml` to pull the image locally
    pull_policy: never
    restart: always
    healthcheck:
      # Docker plugins always run on port 7777
      test: ["CMD", "bash", "-c", "echo -n '' > /dev/tcp/localhost/7777"]
      interval: 5s
      timeout: 30s
      retries: 5
  postgres:
    container_name: postgres
    image: postgres:15
    restart: always
    environment:
      POSTGRES_PASSWORD: pass
    ports:
      - "5432:5432"
    healthcheck:
      test: ["CMD", "pg_isready", "-U", "postgres"]
      interval: 5s
      timeout: 30s
      retries: 5
configs:
  spec.yml:
    content: |
      kind: source
      spec:
        name: airtable
        registry: grpc
        # Notice we use the container name as the host to connect via Docker internal DNS
        path: airtable:7777
        tables: ["*"]
        destinations: ["postgresql"]
        spec:
          access_token: "${AIRTABLE_ACCESS_TOKEN}"
      ---
      kind: destination
      spec:
        name: "postgresql"
        path: "cloudquery/postgresql"
        version: "VERSION_DESTINATION_POSTGRESQL"
        spec:
          # Notice we use the container name as the host to connect via Docker internal DNS
          connection_string: "postgresql://postgres:pass@postgres:5432/postgres?sslmode=disable"
```
3. Run `CLOUDQUERY_API_KEY=<cloudquery-api-key> AIRTABLE_ACCESS_TOKEN=<airtable-access-token> docker compose up -d`
4. You can check the logs of the CLI container to see the sync process. Run `docker logs -f cli` or `docker logs -f airtable` to see the logs.
5. To see the results, you can connect to the PostgreSQL database using your favorite client, for example, `psql`.

## Cleanup

Run `docker compose down` to stop and remove the containers.
