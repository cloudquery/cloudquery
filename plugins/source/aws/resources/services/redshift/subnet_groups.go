package redshift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SubnetGroups() *schema.Table {
	tableName := "aws_redshift_subnet_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_ClusterSubnetGroup.html`,
		Resolver:    fetchRedshiftSubnetGroups,
		Transform:   transformers.TransformWithStruct(&types.ClusterSubnetGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "arn",
				Type:        schema.TypeString,
				Resolver:    resolveSubnetGroupArn(),
				Description: `The Amazon Resource Name (ARN) for the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchRedshiftSubnetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config redshift.DescribeClusterSubnetGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().Redshift
	for {
		response, err := svc.DescribeClusterSubnetGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.ClusterSubnetGroups
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func resolveSubnetGroupArn() schema.ColumnResolver {
	return client.ResolveARN(client.RedshiftService, func(resource *schema.Resource) ([]string, error) {
		return []string{fmt.Sprintf("subnetgroup:%s", *resource.Item.(types.ClusterSubnetGroup).ClusterSubnetGroupName)}, nil
	})
}
