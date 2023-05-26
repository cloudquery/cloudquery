package monitors

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

func MonitorDowntimes() *schema.Table {
	return &schema.Table{
		Name:     "datadog_monitor_downtimes",
		Resolver: fetchMonitorDowntimes,
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveAccountName,
			},
			{
				Name:     "active",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("Active"),
			},
			{
				Name:     "active_child",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("ActiveChild"),
			},
			{
				Name:     "canceled",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Canceled"),
			},
			{
				Name:     "creator_id",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("CreatorId"),
			},
			{
				Name:     "disabled",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("Disabled"),
			},
			{
				Name:     "downtime_type",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("DowntimeType"),
			},
			{
				Name:     "end",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("End"),
			},
			{
				Name:     "id",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "message",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Message"),
			},
			{
				Name:     "monitor_id",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("MonitorId"),
			},
			{
				Name:     "monitor_tags",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: schema.PathResolver("MonitorTags"),
			},
			{
				Name:     "mute_first_recovery_notification",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("MuteFirstRecoveryNotification"),
			},
			{
				Name:     "parent_id",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("ParentId"),
			},
			{
				Name:     "recurrence",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Recurrence"),
			},
			{
				Name:     "scope",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: schema.PathResolver("Scope"),
			},
			{
				Name:     "start",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("Start"),
			},
			{
				Name:     "timezone",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Timezone"),
			},
			{
				Name:     "updater_id",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("UpdaterId"),
			},
		},
	}
}

func fetchMonitorDowntimes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(datadogV1.Monitor)
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.DowntimesAPI.ListMonitorDowntimes(ctx, *p.Id)
	if err != nil {
		return err
	}
	res <- resp
	return nil
}
