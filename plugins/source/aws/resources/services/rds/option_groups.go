package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func OptionGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_option_groups",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_OptionGroup.html`,
		Resolver:    fetchOptionGroups,
		Transform:   transformers.TransformWithStruct(&types.OptionGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OptionGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchOptionGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Rds
	config := rds.DescribeOptionGroupsInput{}
	p := rds.NewDescribeOptionGroupsPaginator(svc, &config)
	for p.HasMorePages() {
		page, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.OptionGroupsList
	}
	return nil
}
