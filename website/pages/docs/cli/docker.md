---
title: Docker
---

import { Callout } from 'nextra-theme-docs'

# Docker

It is possible to use CloudQuery in an isolated container, you can pull the relevant image with the docker commands shown on [getting started](/docs/getting-started/getting-started-with-aws) guide.

## Configuration

CloudQuery uses a YAML file as the primary means of configuration, you can execute the [`cloudquery init`](/docs/cli/commands/cloudquery_init) to generate a file if you do not already have one. For the CloudQuery docker container to use this configuration file you will need to mount the volume to the container like so:

```docker
docker run \
  -v <ABSOLUTE PATH TO CONFIG>/cloudquery.yml:/config/cloudquery.yml \
  ghcr.io/cloudquery/cloudquery:latest \
  fetch --config /config/cloudquery.yml
```

As with running any `cloudquery` command on your CLI you can override the config with the [optional flags](/docs/cli/commands/cloudquery_options) with the docker container. You will also need to make sure you load any ENV variables for providers, such as your `AWS_*` keys etc.

<Callout type="info">

If you are running Docker on an ARM Apple device and you see a segmentation fault when running the container like so `qemu: uncaught target signal 11 (Segmentation fault) - core dumped`; please make sure you are running the latest Docker for Mac release.

</Callout>

## Caching

Due to the way `cloudquery` is [architected](/docs/developers/architecture) it downloads all the components to interact with providers and policies. This means that with a docker container it runs the download step each state as the local cache is lost between executions. To avoid this we recommend mounting a volume to cache the data and configuring `cloudquery` to use this via the `--data-dir` optional flag. An example of this would be:

```docker
docker run \
  -v <PATH TO CACHE>/.cq:/cache/.cq \
  -v <PATH TO CONFIG>/cloudquery.yml:/config/cloudquery.yml \
  ghcr.io/cloudquery/cloudquery:latest \
  fetch --config /config/cloudquery.yml \
    --data-dir /cache/.cq
```

<Callout type="info">

Depending on your operating system, the built components maybe different between your local system and the container. To avoid the different please use a separate cache directory for the container than a local instance of `cloudquery`.

</Callout>

## Fetching data

For the specifics of how `cloudquery fetch` works, and what additional flags it uses please consult the [command page](/docs/cli/commands/cloudquery_fetch). The command will check the `cloudquery.yml` provided and proceed to download any providers defined, uploading the schema to the database, and retrieving the service data from the provider. An example of this command using it with AWS would be:

```docker
docker run \
  -e AWS_ACCESS_KEY_ID=<YOUR AWS ACCESS KEY ID> \
  -e AWS_SECRET_ACCESS_KEY=<YOU AWS SECRET ACCESS KEY> \
  -v ~/Development/cloudquery-grafana/cloudquery.yml:/config/cloudquery.yml \
  ghcr.io/cloudquery/cloudquery:latest \
  fetch --config /config/cloudquery.yml
```

<Callout type="info">

The docker container is set to be verbose by default, which produces a lot of console logs. This is by design.

</Callout>
