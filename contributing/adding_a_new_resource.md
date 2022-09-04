# Adding a New Resource

This guide will help you add a new resource to an existing source plugin (a.k.a. "provider", such as AWS, GCP, Azure or K8s). If you wish to support a completely new cloud platform, first see [creating a new plugin](./creating_a_new_plugin.md). 

## General Guidance

## Installing the tools

 - Install `mockgen` and `cq-gen` by running `make install-tools` in the root of this repository.

## Setting up the service

If the service to which the resource belongs has not been used before in the plugin, there are a few steps that need to be done to configure it.

1. Create the service interface in `client/services.go`
   * Don't forget to add the new service interface name to the go:generate comment.
1. Add the service to the `Services` struct in the `client/client.go`
1. Init the service in the `initServices` function in `client/client.go`
1. Run `go generate client/services.go` to create a mock for your new service. This will update `client/mocks/mock_\<service\>.go` automatically

> If you get an error about not being able to find `mockgen`, run `make install-tools` to install it. If it still fails, run `export PATH=${PATH}:`go env GOPATH`/bin` in you shell to set up your `PATH` environment properly

> You might need to update an existing client by running `go get <path to client>@latest` and then `go mod tidy`

## Setting up the resource

For most resources, we use our open-source tool, `cq-gen`, to generate the code and documentation from a source SDK. For example,
cq-gen can be configured to point to a specific Go struct. It will then recursively read all the fields and comments on this struct
and generate the necessary structures and transformations to load it into a target database. This configuration is done via an `.hcl` config file,
and its structure is documented in the [cq-gen repository](https://github.com/cloudquery/cq-gen).

The only code that needs to be written by you, the human, are the SDK calls to list or describe the resources. Such glue functions
are called "resolvers".

Here are the general steps to follow:
 - Find an appropriate Go SDK function (or OpenAPI definition) that fetches the resource you are interested in. 
 - Note the name of the return type that contains the information you want to read. This will be passed to the cq-gen config via the `path` parameter.
 - Create a new directory for your resource under `resources/services`
 - Define the initial cq-gen `gen.hcl` config inside this directory. For this step it is useful to reference configs of other resources in the same plugin/provider. These will be `.hcl` files contained within the service directory.
 - Run `cq-gen --config gen.hcl --output .`
 - cq-gen will leave some functions unimplemented. You may fill in these function bodies with calls to the appropriate SDK functions.
 - To regenerate from updated config, either run the same command again, or use `go generate`

## Specific Guides

See the following guides for deep-dives into adding resources for specific source plugins:
 - [Adding a new resource (AWS)](../plugins/source/aws/docs/contributing/adding_a_new_resource.md)
 - [Adding a new resource (Azure)](../plugins/source/azure/docs/contributing/adding_a_new_resource.md)
 - [Adding a new resource (GCP)](../plugins/source/gcp/docs/contributing/adding_a_new_resource.md)
