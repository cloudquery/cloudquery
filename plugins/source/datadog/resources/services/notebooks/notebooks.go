package notebooks

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Notebooks() *schema.Table {
	return &schema.Table{
		Name:      "datadog_notebooks",
		Resolver:  fetchNotebooks,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV1.NotebooksResponseData{}, transformers.WithPrimaryKeys("Id")),
		Columns:   schema.ColumnList{client.AccountNameColumn},
	}
}

func fetchNotebooks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.NotebooksAPI.ListNotebooks(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
