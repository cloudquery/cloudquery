package notebooks

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Notebooks() *schema.Table {
	return &schema.Table{
		Name:      "datadog_notebooks",
		Resolver:  fetchNotebooks,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV1.NotebooksResponseData{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			client.AccountNameColumn,
		},
	}
}

func fetchNotebooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.NotebooksAPI.ListNotebooks(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
