package slos

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Corrections() *schema.Table {
	return &schema.Table{
		Name:      "datadog_slo_corrections",
		Resolver:  fetchCorrections,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV1.SLOCorrection{}, transformers.WithPrimaryKeys("Id")),
		Columns:   schema.ColumnList{client.AccountNameColumn},
	}
}

func fetchCorrections(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV2(ctx)
	resp, _, err := c.DDServices.ServiceLevelObjectiveCorrectionsAPI.ListSLOCorrection(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
