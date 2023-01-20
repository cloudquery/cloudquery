package export

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/pkg/errors"
)

func Events() *schema.Table {
	return &schema.Table{
		Name:      "mixpanel_export_events",
		Resolver:  fetchExportEvents,
		Transform: transformers.TransformWithStruct(&mixpanel.ExportEvent{}, client.SharedTransformers()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeInt,
				Resolver: client.ResolveProjectID,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "time",
				Type:     schema.TypeTimestamp,
				Resolver: resolveExportTime,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchExportEvents(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	ret, err := cl.Services.ExportData(ctx, cl.MPSpec.StartDate, cl.MPSpec.EndDate)
	if err != nil {
		return err
	}
	res <- ret
	return nil
}

func resolveExportTime(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	e := r.Item.(mixpanel.ExportEvent)
	ts, ok := e.Properties["time"]
	if !ok {
		return errors.New("event does not have a time property")
	}
	tf, ok := ts.(float64) // json.Number translates to float64
	if !ok {
		return fmt.Errorf("event time property is not a float: %T", ts)
	}
	return r.Set(c.Name, time.Unix(int64(tf), 0))
}
