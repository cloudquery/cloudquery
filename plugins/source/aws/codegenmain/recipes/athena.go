package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var AthenaResources = []*Resource{
	{
		DefaultColumns:       []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:            &types.DataCatalog{},
		AWSService:           "Athena",
		AWSSubService:        "DataCatalogs",
		Template:             "resource_list_and_detail",
		ListVerb:             "List",
		ListFieldName:        "DataCatalogsSummary",
		ResponseItemsName:    "CatalogName",
		ItemName:             "DataCatalog",
		DetailInputFieldName: "Name",
		ResponseItemsType:    "DataCatalogSummary",
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		CustomErrorBlock: `
		// retrieving of default data catalog (AwsDataCatalog) returns "not found error" but it exists and its
		// relations can be fetched by its name
		if *itemSummary.CatalogName == "AwsDataCatalog" {
			resultsChan <- types.DataCatalog{Name: itemSummary.CatalogName, Type: itemSummary.Type}
			return
		}
`,
		CustomTagField: `aws.String(createDataCatalogArn(cl, *item.Name))`,
		ColumnOverrides: map[string]codegen.ColumnDefinition{
			"arn": {
				Type:     schema.TypeString,
				Resolver: "resolveAthenaDataCatalogArn",
			},
			"tags": {
				Type:        schema.TypeJSON,
				Description: "Tags associated with the Athena data catalog.",
				Resolver:    ResolverAuto,
			},
		},
		CustomResolvers: []string{
			`
func resolveAthenaDataCatalogArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	dc := resource.Item.(types.DataCatalog)
	return diag.WrapError(resource.Set(c.Name, createDataCatalogArn(cl, *dc.Name)))
}

func createDataCatalogArn(cl *client.Client, catalogName string) string {
	return cl.ARN(client.Athena, "datacatalog", catalogName)
}
`,
		},
	},
}
