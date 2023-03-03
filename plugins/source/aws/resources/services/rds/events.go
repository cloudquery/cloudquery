package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Events() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_events",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DescribeEvents.html`,
		Resolver:    fetchRdsEvents,
		Transform:   transformers.TransformWithStruct(&types.Event{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchRdsEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Rds
	duration := int32(60 * 24 * 14) // 14 days (maximum)
	config := rds.DescribeEventsInput{
		Duration: &duration,
	}
	p := rds.NewDescribeEventsPaginator(svc, &config)
	for p.HasMorePages() {
		page, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Events
	}
	return nil
}
