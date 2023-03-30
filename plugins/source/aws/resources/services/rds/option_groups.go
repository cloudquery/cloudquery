package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func OptionGroups() *schema.Table {
	tableName := "aws_rds_option_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_OptionGroup.html`,
		Resolver:    fetchOptionGroups,
		Transform:   client.TransformWithStruct(&types.OptionGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
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
