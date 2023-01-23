package costmanagement

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func view_queries() *schema.Table {
	return &schema.Table{
		Name:        "azure_costmanagement_view_queries",
		Resolver:    fetchViewQueries,
		Description: "https://learn.microsoft.com/en-us/rest/api/cost-management/query/usage?tabs=HTTP#queryresult",
		Transform:   transformers.TransformWithStruct(&armcostmanagement.QueryResult{}, transformers.WithSkipFields("ETag")),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ETag"),
			},
		},
	}
}
