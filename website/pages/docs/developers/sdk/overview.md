# Overview

**The CloudQuery SDK** enables building CloudQuery providers which allows CloudQuery's users to **E**xtract/**T**ransform/**L**oad existing and popular service provider APIs as well as custom in-house solutions into a SQL database.
It was developed to allow easy creation for providers and reduce boilerplate code when developing them, enabling developers to solely focus on the **E**xtract while the SDK takes care of the rest.

The core idea of the SDK is to allow developing a provider in a straightforward manner, while allowing the user free control in fetching the data to be inserted.

## Key Concepts

- **Provider** structs are the core component of the SDK that require the implementor to only set a list all available resources, configuration of it's client and it's config. The provider structs implements the CQProvider Interface allowing the user to just implement his resources and configure function.
- **Table** is the main building block in the SDK provider schema, these tables are passed to the Provider to define what resources the provider supports. [Tables](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/table.go) define their columns, relations (which are also tables). Each table has a resolver function that is called by the SDK with the client that was configured early by the user implementation.
- **Resolvers** are functions to fetch resource metadata from the cloud provider, look up values inside structs, or simply do conversion between them.

## Getting Started

Take a look at the [CloudQuery Architecture](../architecture) and [Developing a New Provider](../developing-new-provider) documents.

## Debug and Test

To debug a provider see the [Debugging](../debugging) section.

The SDK also provides a [testing](./testing) package/helper to minimize the boilerplate needed to test your provider.
