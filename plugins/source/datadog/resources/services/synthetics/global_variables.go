package synthetics

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func GlobalVariables() *schema.Table {
	return &schema.Table{
		Name:      "datadog_global_variables",
		Resolver:  fetchGlobalVariables,
		Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&datadogV1.SyntheticsGlobalVariable{}),
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccountName,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
