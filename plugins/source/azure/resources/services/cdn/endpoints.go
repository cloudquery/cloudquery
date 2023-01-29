package cdn

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func endpoints() *schema.Table {
	return &schema.Table{
		Name:        "azure_cdn_endpoints",
		Resolver:    fetchEndpoints,
		Description: "https://learn.microsoft.com/en-us/rest/api/cdn/endpoints/list-by-profile?tabs=HTTP#endpoint",
		Transform:   transformers.TransformWithStruct(&armcdn.Endpoint{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{},
	}
}
