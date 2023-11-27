---
title: Building CloudQuery
description: Learn how to build CloudQuery from source.
---

# Building From Source

The preferred way to use CloudQuery is through the available distribution, see the [Quickstart](/docs/quickstart) section of the [Docs](/docs)

To build CloudQuery CLI from source, follow the steps:

1. CloudQuery is developed in Go. Ensure you have a working [Go runtime](https://go.dev/)
2. Fork and clone the CloudQuery repository. If you’re not sure how to do this, please watch [these videos](https://egghead.io/courses/how-to-contribute-to-an-open-source-project-on-github).
3. From the cloned repository root, change directory to `./cli` and run `go build -o cloudquery` to build the CloudQuery CLI. The binary will be created in the same directory.

Building a plugin from source is similar. Most plugins have a makefile in their directory to make this easier. For example, to build the `aws` plugin, run `make build` from the `./plugins/source/aws` directory. The resulting binary can be used by providing the path to it as the `path` parameter in a [plugin config](/docs/reference/source-spec), together with the `local` registry. Python plugins have `make build-docker` to build a Docker image that can be used with the `docker` registry.


