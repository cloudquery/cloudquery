---
title: Docker
---

import { Callout } from 'nextra-theme-docs'

# Docker

It is possible to use CloudQuery in an isolated container, you can pull the relevant image via `docker pull ghcr.io/cloudquery/cloudquery:latest`.

## Configuration

CloudQuery uses YAML files as the primary means of configuration. For the CloudQuery docker container to use this configuration file you will need to mount the volume to the container like so:

```docker copy
docker run \
  -v <ABSOLUTE_PATH_TO_CONFIG_FILE>:/config.yml \
  # set any env variable with -e <ENV_VAR_NAME>=<ENV_VAR_VALUE>
  ghcr.io/cloudquery/cloudquery:latest \
  sync /config.yml
```

As with running any `cloudquery` command on your CLI you can override the config with the [optional flags](/docs/reference/cli/cloudquery) with the docker container. You will also need to make sure you load any ENV variables for source and destination plugins, such as your `AWS_*` keys etc.

<Callout type="info">

If you are running Docker on an ARM Apple device and you see a segmentation fault when running the container like so `qemu: uncaught target signal 11 (Segmentation fault) - core dumped`; please make sure you are running the latest Docker for Mac release.

</Callout>

## Caching

Due to the way `cloudquery` is [architected](/docs/developers/architecture) it downloads all the components to interact with source and destination plugins. This means that with a docker container it runs the download step each state as the local cache is lost between executions. To avoid this we recommend mounting a volume to cache the data and configuring `cloudquery` to use this via the `--data-dir` optional flag. An example of this would be:

```docker copy
docker run \
  -v <PATH TO CACHE>/.cq:/cache/.cq \
  -v <PATH TO CONFIG>/cloudquery.yml:/config/cloudquery.yml \
  ghcr.io/cloudquery/cloudquery:latest \
  fetch --config /config/cloudquery.yml \
```

<Callout type="info">

Depending on your operating system, the built components maybe different between your local system and the container. To avoid the different please use a separate cache directory for the container than a local instance of `cloudquery`.

</Callout>

