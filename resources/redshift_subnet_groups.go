package resources

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
		Resolver:     fetchRedshiftSubnetGroups,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "cluster_subnet_group_name",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "subnet_group_status",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRedshiftSubnetGroupTags,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_redshift_subnet_group_subnets",
				Resolver: fetchRedshiftSubnetGroupSubnets,
				Columns: []schema.Column{
					{
						Name:     "subnet_group_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "subnet_availability_zone_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetAvailabilityZone.Name"),
					},
					{
						Name:     "subnet_availability_zone_supported_platforms",
						Type:     schema.TypeStringArray,
						Resolver: resolveRedshiftSubnetGroupSubnetSubnetAvailabilityZoneSupportedPlatforms,
					},
					{
						Name: "subnet_identifier",
						Type: schema.TypeString,
					},
					{
						Name: "subnet_status",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRedshiftSubnetGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
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
func resolveRedshiftSubnetGroupTags(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.ClusterSubnetGroup)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchRedshiftSubnetGroupSubnets(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	clusterSubnetGroup, ok := parent.Item.(types.ClusterSubnetGroup)
	if !ok {
		return fmt.Errorf("not redshift cluster subnet group")
	}
	res <- clusterSubnetGroup.Subnets
	return nil
}
func resolveRedshiftSubnetGroupSubnetSubnetAvailabilityZoneSupportedPlatforms(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.Subnet)
	platforms := make([]*string, len(r.SubnetAvailabilityZone.SupportedPlatforms))
	for i, p := range r.SubnetAvailabilityZone.SupportedPlatforms {
		platforms[i] = p.Name
	}
	return resource.Set("subnet_availability_zone_supported_platforms", platforms)
}
