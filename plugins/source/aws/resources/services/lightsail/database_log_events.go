package lightsail

import (
	"context"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lightsail/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"golang.org/x/sync/errgroup"
)

func databaseLogEvents() *schema.Table {
	tableName := "aws_lightsail_database_log_events"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetRelationalDatabaseLogEvents.html`,
		Resolver:    fetchLightsailDatabaseLogEvents,
		Transform:   transformers.TransformWithStruct(&models.LogEventWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "database_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchLightsailDatabaseLogEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.RelationalDatabase)
	input := lightsail.GetRelationalDatabaseLogStreamsInput{
		RelationalDatabaseName: r.Name,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Lightsail
	streams, err := svc.GetRelationalDatabaseLogStreams(ctx, &input, func(options *lightsail.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	endTime := time.Now()
	startTime := endTime.Add(-time.Hour * 24 * 14) // two weeks
	errs, ctx := errgroup.WithContext(ctx)
	errs.SetLimit(MaxGoroutines)
	for _, s := range streams.LogStreams {
		func(database, stream string, startTime, endTime time.Time) {
			errs.Go(func() error {
				return fetchLogEvents(ctx, res, cl, database, stream, startTime, endTime)
			})
		}(*r.Name, s, startTime, endTime)
	}
	err = errs.Wait()
	if err != nil {
		return err
	}
	return nil
}

func fetchLogEvents(ctx context.Context, res chan<- any, cl *client.Client, database, stream string, startTime, endTime time.Time) error {
	svc := cl.Services().Lightsail
	input := lightsail.GetRelationalDatabaseLogEventsInput{
		RelationalDatabaseName: &database,
		LogStreamName:          &stream,
		StartTime:              &startTime,
		EndTime:                &endTime,
	}
	// No paginator available
	for {
		response, err := svc.GetRelationalDatabaseLogEvents(ctx, &input, func(options *lightsail.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, e := range response.ResourceLogEvents {
			res <- models.LogEventWrapper{
				LogEvent:      e,
				LogStreamName: stream,
			}
		}
		if aws.ToString(response.NextForwardToken) == "" || len(response.ResourceLogEvents) == 0 {
			break
		}
		input.PageToken = response.NextForwardToken
	}
	return nil
}
