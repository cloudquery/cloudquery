package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func RdsSubnetGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_rds_subnet_groups",
		Resolver:     fetchRdsSubnetGroups,
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
				Name:     "db_subnet_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroupArn"),
			},
			{
				Name:     "db_subnet_group_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroupDescription"),
			},
			{
				Name:     "db_subnet_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroupName"),
			},
			{
				Name: "subnet_group_status",
				Type: schema.TypeString,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_rds_subnet_group_subnets",
				Resolver: fetchRdsSubnetGroupSubnets,
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
						Name: "subnet_identifier",
						Type: schema.TypeString,
					},
					{
						Name:     "subnet_outpost_arn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("SubnetOutpost.Arn"),
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
func fetchRdsSubnetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config rds.DescribeDBSubnetGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().RDS
	for {
		response, err := svc.DescribeDBSubnetGroups(ctx, &config, func(o *rds.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.DBSubnetGroups
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func fetchRdsSubnetGroupSubnets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	subnetGroup, ok := parent.Item.(types.DBSubnetGroup)
	if !ok {
		return fmt.Errorf("not db cluster")
	}
	res <- subnetGroup.Subnets
	return nil
}
