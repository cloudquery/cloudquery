package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Inspector2Resources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "findings",
			Struct:      &types.Finding{},
			Description: "https://docs.aws.amazon.com/inspector/v2/APIReference/API_Finding.html",
			SkipFields:  []string{"FindingArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("FindingArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "inspector2"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("inspector2")`
	}
	return resources
}
