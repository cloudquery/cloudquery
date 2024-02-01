package cloudtrail

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

const tableName = "aws_cloudtrail_events"

func Events() *schema.Table {
	return &schema.Table{
		Name:          tableName,
		Description:   `https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_Event.html`,
		Resolver:      fetchCloudtrailEvents,
		Multiplex:     client.ServiceAccountRegionMultiplexer(tableName, "cloudtrail"),
		Transform:     transformers.TransformWithStruct(&types.Event{}, transformers.WithPrimaryKeyComponents("EventId")),
		IsIncremental: true,
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "cloud_trail_event",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("CloudTrailEvent"),
			},
			{
				Name:           "event_time",
				Type:           arrow.FixedWidthTypes.Timestamp_us,
				Resolver:       schema.PathResolver("EventTime"),
				IncrementalKey: true,
			},
		},
	}
}

func fetchCloudtrailEvents(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudtrail).Cloudtrail

	stateClient := cl.StateClient()
	le := cloudtrail.LookupEventsInput{}
	var backendKey string
	// Retrieve the last event time from the backend for this table option config.
	// We use a hash of the config as the key, so changing the config will cause a full refresh.
	value, err := stateClient.GetKey(ctx, tableName+cl.ID())
	if err != nil {
		return fmt.Errorf("failed to retrieve state from backend: %w", err)
	}

	if value != "" {
		date, err := time.Parse(time.RFC3339Nano, value)
		if err != nil {
			return fmt.Errorf("retrieved invalid state value: %q %w", value, err)
		}
		le.StartTime = &date
	}

	var lastEventTime *time.Time
	// var err error
	paginator := cloudtrail.NewLookupEventsPaginator(svc, &le)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudtrail.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Events

		// Retrieve the timestamp from the latest event
		for _, event := range page.Events {
			if lastEventTime == nil {
				lastEventTime = event.EventTime
				continue
			}
			if event.EventTime.After(*lastEventTime) {
				lastEventTime = event.EventTime
			}
		}
	}

	if lastEventTime != nil {
		err := stateClient.SetKey(ctx, tableName+backendKey, lastEventTime.Format(time.RFC3339Nano))
		if err != nil {
			return fmt.Errorf("failed to save state to backend: %w", err)
		}
	}

	return nil
}
