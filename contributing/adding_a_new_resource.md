# Adding a New Resource

This guide will help you add a new resource to an existing source integration (a.k.a. "provider", such as AWS, GCP, Azure or K8s). If you wish to support a completely new cloud platform, first see [creating a new integration](./creating_a_new_integration.md).

## General Guidance

## Setting up the service

If the service to which the resource belongs has not been used before in the integration, there are a few steps that need to be done to configure it.

1. Create the service interface in `client/services.go`
   - Don't forget to add the new service interface name to the go:generate comment.
1. Add the service to the `Services` struct in the `client/client.go`
1. Init the service in the `initServices` function in `client/client.go`
1. Run `go generate client/services.go` to create a mock for your new service. This will update `client/mocks/mock_\<service\>.go` automatically

> If you get an error about not being able to find `mockgen`, run `make install-tools` to install it. If it still fails, run `export PATH=${PATH}:`go env GOPATH`/bin` in you shell to set up your `PATH` environment properly
> You might need to update an existing client by running `go get <path to client>@latest` and then `go mod tidy`

## Setting up the resource

We use code generation to generate the code from a source SDK. This functionality is provided by the CloudQuery plugin-sdk. It will read all the fields on a given struct and generate the necessary structures and transformations to load it into a destination database. The configuration is done via "recipe" files, contained in the `codegen/recipes` directory for each source plugin.

The only code that needs to be written by you are the SDK calls to list or describe the resources. Such glue functions are called "resolvers".

Here are the general steps to follow:

- Find an appropriate Go SDK function that fetches the resource you are interested in.
- Note the type of the return type that contains the information you want to read. This will be passed to `codegen` via the `Struct` property.
- Create a new recipe file for the resource, if one does not exist already.
- Add the resource to the recipe file.
- Run `go run main.go` in the `codegen` directory. The generated table will be in `plugins/source/<plugin>/resources/services/<service>/<resource>.go`.
- To regenerate from updated configuration, re-run `go run main.go` from the `codegen` directory again.
- Implement one or more resolver functions (as referenced by the generated file) in `plugins/source/<plugin>/resources/services/<service>/<resource>_fetch.go`.
- Add a mock test for the resource in `plugins/source/<plugin>/resources/services/<service>/<resource>_mock_test.go`

## Specific Guides

See the following guides for deep-dives into adding resources for specific source integrations:

- [Adding a new resource (AWS)](../plugins/source/aws/CONTRIBUTING.md)
- [Adding a new resource (Azure)](../plugins/source/azure/CONTRIBUTING.md)
- [Adding a new resource (GCP)](../plugins/source/gcp/CONTRIBUTING.md)
