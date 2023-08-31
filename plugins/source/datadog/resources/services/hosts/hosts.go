package hosts

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Hosts() *schema.Table {
	return &schema.Table{
		Name:      "datadog_hosts",
		Resolver:  fetchHosts,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV1.Host{}, transformers.WithPrimaryKeys("Id")),
		Columns:   schema.ColumnList{client.AccountNameColumn},
	}
}

func fetchHosts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.HostsAPI.ListHosts(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetHostList()
	return nil
}
