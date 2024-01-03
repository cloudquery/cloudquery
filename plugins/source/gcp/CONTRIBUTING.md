# GCP Source Plugin Contribution Guide


***Please note, GCP is `open core` , before you begin adding any resources to this plugin,please file an issue before opening a PR. Not all contributions will be accepted especially if they are part of the commercial plugin offering***

This document serves as a guide for adding new services and resources to the GCP source plugin.

In the following steps, we will use the fictional `MyService` GCP service with `MyResource` resource as an example.


##  Adding a new resource for a new service


1. Create a folder in the `resources/services` directory with the name of the service. For example, if the service is called `MyService`, create a folder called `myservice`.
2. Create a file in the folder called `myresource.go`. The file should contain a single exported function that is the name of the resource and will return a ` *schema.Table`
3. Specify values for the following struct members:
   `name`: The name of the table. This will be in the form `gcp` 
   `multiplexer`: options include `client.OrgMultiplex`, `client.FolderMultiplex`, `client.ProjectMultiplex`, `client.ProjectMultiplexEnabledServices(serviceDNS string)` For details on each one, see the [multiplexer section](#choosing-a-multiplexer)
   `description`: a short description of the resource, usually will a link to the documentation
   `columns`: an array of columns that will be added in addition to the columns from the data returned by the API. This is typically used to add primary keys if the API does not return them.
4. Create `Resolver` function that will actually resolve the resource.
5. Finally, implement a mock test in `myresource_mock_test.go`.
6. Finally add the top level resource to the `resource/plugin/tables.go` in the `tables` list

We recommend looking at other resources similar to yours to get an idea of what needs to be done in any of the above steps. 

### Implementing Resolver Functions

A few important things to note when adding functions that call the GCP API:

- If possible, always use an API call that allows you to fetch many resources at once
- Take pagination into account. Ensure you fetch **all** the resources.
- Columns may also have their own resolver functions (not covered in this guide). This may be used for simple transformations or when additional calls can help add further context to the table.

## General Tips

- Keep transformations to a minimum. As far as possible, we aim to deliver an accurate reflection of what the API provides.
- We generally only unroll structs one level deep. Nested structs should be transformed into JSON columns.
- If you get stuck or need help, feel free to reach out on [Discord](https://www.cloudquery.io/discord). We are a friendly community and would love to help!


### Choosing a Multiplexer

In the GCP plugin there are three types of multiplexers. Every top level resource needs to use multiplexer:

1. `ProjectMultiplex` (_default_): This is the most basic of multiplexers in that it will resolve the resource in each project that is being synced. 
2. `ProjectMultiplexEnabledServices(serviceDNS string)`:  This multiplexer will only attempt to sync a resource if that project has the service enabled otherwise the resource will be skipped for that specific projectID. On top of this the user must also enable the feature via `enabled_services_only: true` in the spec. In order to use this multiplexer you must specify a valid `resource.ServiceDNS`
3. `client.OrgMultiplex`: For resources that are unique across an entire Organization. In order to use this multiplexer you have to explicitly set the multiplexer `client.OrgMultiplex`