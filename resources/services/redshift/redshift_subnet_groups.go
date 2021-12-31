package redshift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func RedshiftSubnetGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_redshift_subnet_groups",
		Description:  "Describes a subnet group.",
		Resolver:     fetchRedshiftSubnetGroups,
		Multiplex:    client.ServiceAccountRegionMultiplexer("redshift"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "cluster_subnet_group_name"}},
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
				Resolver:    resolveRedshiftSubnetGroupTags,
			},
			{
				Name:        "vpc_id",
				Description: "The VPC ID of the cluster subnet group.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_redshift_subnet_group_subnets",
				Description: "Describes a subnet.",
				Resolver:    fetchRedshiftSubnetGroupSubnets,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"subnet_group_cq_id", "subnet_identifier"}},
				Columns: []schema.Column{
					{
						Name:        "subnet_group_cq_id",
						Description: "Unique CloudQuery ID of aws_redshift_subnet_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "subnet_availability_zone_name",
						Description: "The name of the availability zone.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetAvailabilityZone.Name"),
					},
					{
						Name:        "subnet_availability_zone_supported_platforms",
						Description: "A list of supported platforms for orderable clusters.",
						Type:        schema.TypeStringArray,
						Resolver:    resolveRedshiftSubnetGroupSubnetSubnetAvailabilityZoneSupportedPlatforms,
					},
					{
						Name:        "subnet_identifier",
						Description: "The identifier of the subnet.",
						Type:        schema.TypeString,
					},
					{
						Name:        "subnet_status",
						Description: "The status of the subnet.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRedshiftSubnetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config redshift.DescribeClusterSubnetGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().Redshift
	for {
		response, err := svc.DescribeClusterSubnetGroups(ctx, &config, func(o *redshift.Options) {
			o.Region = c.Region
		})
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
func resolveRedshiftSubnetGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ClusterSubnetGroup)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchRedshiftSubnetGroupSubnets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	clusterSubnetGroup, ok := parent.Item.(types.ClusterSubnetGroup)
	if !ok {
		return fmt.Errorf("not redshift cluster subnet group")
	}
	res <- clusterSubnetGroup.Subnets
	return nil
}
func resolveRedshiftSubnetGroupSubnetSubnetAvailabilityZoneSupportedPlatforms(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Subnet)
	platforms := make([]*string, len(r.SubnetAvailabilityZone.SupportedPlatforms))
	for i, p := range r.SubnetAvailabilityZone.SupportedPlatforms {
		platforms[i] = p.Name
	}
	return resource.Set("subnet_availability_zone_supported_platforms", platforms)
}
