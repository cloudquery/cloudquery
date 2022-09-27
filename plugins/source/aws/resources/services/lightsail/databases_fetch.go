package lightsail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lightsail/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
)

func fetchLightsailDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetRelationalDatabasesInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetRelationalDatabases(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.RelationalDatabases
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func fetchLightsailDatabaseParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RelationalDatabase)
	input := lightsail.GetRelationalDatabaseParametersInput{
		RelationalDatabaseName: r.Name,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetRelationalDatabaseParameters(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Parameters
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func fetchLightsailDatabaseEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RelationalDatabase)
	input := lightsail.GetRelationalDatabaseEventsInput{
		RelationalDatabaseName: r.Name,
		DurationInMinutes:      aws.Int32(20160), //two weeks
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetRelationalDatabaseEvents(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.RelationalDatabaseEvents
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func fetchLightsailDatabaseLogEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RelationalDatabase)
	input := lightsail.GetRelationalDatabaseLogStreamsInput{
		RelationalDatabaseName: r.Name,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	streams, err := svc.GetRelationalDatabaseLogStreams(ctx, &input)
	if err != nil {
		return err
	}
	endTime := time.Now()
	startTime := endTime.Add(-time.Hour * 24 * 14) //two weeks
	errs, ctx := errgroup.WithContext(ctx)
	errs.SetLimit(MaxGoroutines)
	for _, s := range streams.LogStreams {
		func(database, stream string, startTime, endTime time.Time) {
			errs.Go(func() error {
				return fetchLogEvents(ctx, res, c, database, stream, startTime, endTime)
			})
		}(*r.Name, s, startTime, endTime)
	}
	err = errs.Wait()
	if err != nil {
		return err
	}
	return nil
}

func fetchLogEvents(ctx context.Context, res chan<- interface{}, c *client.Client, database, stream string, startTime, endTime time.Time) error {
	svc := c.Services().Lightsail
	input := lightsail.GetRelationalDatabaseLogEventsInput{
		RelationalDatabaseName: &database,
		LogStreamName:          &stream,
		StartTime:              &startTime,
		EndTime:                &endTime,
	}
	for {
		response, err := svc.GetRelationalDatabaseLogEvents(ctx, &input)
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
