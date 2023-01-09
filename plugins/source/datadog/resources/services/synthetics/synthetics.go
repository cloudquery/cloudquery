package synthetics

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Synthetics() *schema.Table {
	return &schema.Table{
		Name:      "datadog_synthetics",
		Resolver:  fetchSynthetics,
		Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&datadogV1.SyntheticsTestDetails{}),
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccountName,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "public_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
