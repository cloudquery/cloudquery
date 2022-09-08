package redshift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RedshiftSubnetGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_redshift_subnet_groups",
		Description: "Describes a subnet group.",
		Resolver:    fetchRedshiftSubnetGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("redshift"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.RedshiftService, func(resource *schema.Resource) ([]string, error) {
					return []string{fmt.Sprintf("subnetgroup:%s", *resource.Item.(types.ClusterSubnetGroup).ClusterSubnetGroupName)}, nil
				}),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "cluster_subnet_group_name",
				Description: "The name of the cluster subnet group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "The description of the cluster subnet group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_group_status",
				Description: "The status of the cluster subnet group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The list of tags for the cluster subnet group.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "vpc_id",
				Description: "The VPC ID of the cluster subnet group.",
				Type:        schema.TypeString,
			},
			{
				Name: "subnets",
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchRedshiftSubnetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
