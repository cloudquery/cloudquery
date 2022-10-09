package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SecretsManagerResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "secrets",
			Struct:              &secretsmanager.DescribeSecretOutput{},
			SkipFields:          []string{"ARN", "Tags", "ResultMetadata"},
			PreResourceResolver: "getSecret",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "The list of user-defined tags associated with the secret",
						Type:        schema.TypeJSON,
						Resolver:    `client.ResolveTags`,
					},
					{
						Name:        "policy",
						Description: "A JSON-formatted string that describes the permissions that are associated with the attached secret.",
						Type:        schema.TypeJSON,
						Resolver:    `fetchSecretsmanagerSecretPolicy`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "secretsmanager"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("secretsmanager")`
	}
	return resources
}
