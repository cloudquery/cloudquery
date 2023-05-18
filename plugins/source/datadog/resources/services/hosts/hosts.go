package hosts

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Hosts() *schema.Table {
	return &schema.Table{
		Name:      "datadog_hosts",
		Resolver:  fetchHosts,
		Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&datadogV1.Host{}),
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
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchHosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.HostsAPI.ListHosts(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetHostList()
	return nil
}
