package export

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/pkg/errors"
)

const key = "mixpanel_export_events"

func ExportEvents() *schema.Table {
	return &schema.Table{
		Name:                 "mixpanel_export_events",
		Description:          `https://developer.mixpanel.com/reference/raw-event-export`,
		Resolver:             fetchExportEvents,
		PostResourceResolver: postExportEvents,
		Transform:            transformers.TransformWithStruct(&mixpanel.ExportEvent{}, client.SharedTransformers(transformers.WithPrimaryKeys("Event"))...),
		IsIncremental:        true,
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
					PrimaryKey:     true,
					IncrementalKey: true,
				},
			},
			{
				Name:     "distinct_id",
				Type:     schema.TypeString,
				Resolver: resolveDistinctID,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchExportEvents(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	cursor := int64(0)
	if cl.Backend != nil {
		value, err := cl.Backend.Get(ctx, key, cl.ID())
		if err != nil {
			return fmt.Errorf("failed to retrieve state from backend: %w", err)
		}
		if value != "" {
			valInt, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("retrieved invalid state value: %q %w", value, err)
			}
			cursor = valInt
		}
	}

	err := cl.Services.ExportData(ctx, cl.MPSpec.StartDate, cl.MPSpec.EndDate, cursor, res)
	return err
}

func postExportEvents(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	if cl.Backend == nil {
		return nil
	}

	ev := resource.Item.(mixpanel.ExportEvent)
	ts, ok := ev.Properties["time"].(float64)
	if !ok {
		cl.Logger().Warn().Msg("postExportEvents: event does not have a time property") // shouldn't happen as resolveExportTime would error out first
		return nil
	}

	if err := cl.Backend.SetHWM(ctx, key, cl.ID(), int64(ts)); err != nil {
		return fmt.Errorf("failed to store state in backend: %w", err)
	}

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

func resolveDistinctID(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	e := r.Item.(mixpanel.ExportEvent)
	val, ok := e.Properties["distinct_id"]
	if !ok {
		return errors.New("event does not have a distinct_id property")
	}
	return r.Set(c.Name, val)
}
