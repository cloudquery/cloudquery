package slos

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Objectives() *schema.Table {
	return &schema.Table{
		Name:      "datadog_slos",
		Resolver:  fetchObjectives,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV1.ServiceLevelObjective{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			client.AccountNameColumn,
		},
	}
}

func fetchObjectives(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV2(ctx)
	resp, _, err := c.DDServices.ServiceLevelObjectivesAPI.ListSLOs(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
