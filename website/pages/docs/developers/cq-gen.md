# Generating resources

Adding resources to a provider can sometimes be a tedious task, some resources can have more than hundreds of fields and relations, and adding them all can
take a long time. To remedy this issue, the [cq-gen](https://github.com/cloudquery/cq-gen) project was created. cq-gen allows to easily generate more of the boilerplate code for resources from common specs such as go code, OpenAPI specs, protobuf, and graphql.

## Set up cq-gen in your provider project

If you haven't created a provider use this command to create a project or alternatively use our [template](https://github.com/cloudquery/cq-provider-template) repository as a base.

```bash
mkdir cq-my-provider
cd cq-my-provider

go mod init github.com/[username]/[cq-my-provider]

# install the `cq-gen` binary under your go path
go install github.com/cloudquery/cq-gen
```

Next, create a tools.go file and add cq-gen as a tool dependency for your module.

```go
//go:build tools

package tools

import (
	_ "github.com/cloudquery/cq-gen"
)
```

To automatically add the dependency to your go.mod run:

```bash
go mod tidy
```

## Running cq-gen for the first time

To run `cq-gen` you must first create a resource hcl config. As an example we will create the AWS CloudFormation [stacks](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/cloudformation@v1.20.0/types#Stack)

```hcl
service = "aws"
output_directory = "."
add_generate = true

resource "aws" "cloudformation" "stacks" {
  path = "github.com/aws/aws-sdk-go-v2/service/cloudformation/types.Stack"
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreAccessDeniedServiceDisabled"
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cq-provider-aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = ["id"]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    type        = "string"
    description = "The AWS Region of the resource."
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion"
    }
  }

}
```

\*Note: we will call this file stacks.hcl

To execute this configuration we will run our cq-gen tool with the following command:

```bash
# cq-gen has to run from the directory of the `gen.hcl` file
cd resources/services/<service-name>
cq-gen --resource stacks --domain cloudformation --config gen.hcl
```

The command above will generate the following code:

```go
package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource stacks --config gen.hcl --output .
func Stacks() *schema.Table {
	return &schema.Table{
		Name:         "aws_cloudformation_stacks",
		Description:  "The Stack data type.",
		Resolver:     fetchCloudformationStacks,
		Multiplex:    client.ServiceAccountRegionMultiplexer("cloudformation"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
                		Resolver:    client.ResolveAWSRegion,
			},
            ... # shortened for sake for readability
```

As you can see above the function definition, cq-gen, added a `//go:generate` command to help regenerate easily with `go generate`.

## Flags

The cq-gen command is used to generate `schema.Table` from given source (go, protobuf, OpenAPI, graphql, etc') It supports the following flags:

- `-output`: which directory to write the resulting source code.
- `-resource`: the name of the resource to generate as defined in the hcl.
- `-domain`: the domain of the resource to generate as defined in the hcl.
- `-config`: the path to the configuration file of cq-gen on how to generate the resource.

## Configuration

For full documentation on how to use cq-gen and all the available configuration options please check [the repository](https://github.com/cloudquery/cq-gen#configuration).
