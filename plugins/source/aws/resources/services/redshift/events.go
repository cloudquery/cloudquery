package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Events() *schema.Table {
	tableName := "aws_redshift_events"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_Event.html.

Only events occurred in the last hour are returned.`,
		Resolver:  fetchEvents,
		Transform: transformers.TransformWithStruct(&types.Event{}),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Redshift

	config := redshift.DescribeEventsInput{
		// Default values are here for clarity
		Duration:   aws.Int32(60),
		MaxRecords: aws.Int32(100),
	}
	for {
		result, err := svc.DescribeEvents(ctx, &config)
		if err != nil {
			return err
		}
		res <- result.Events
		if aws.ToString(result.Marker) == "" {
			break
		}
		config.Marker = result.Marker
	}
	return nil
}
