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
		Description:  "Contains the details of an Amazon RDS DB subnet group",
		Resolver:     fetchRdsSubnetGroups,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name:        "db_subnet_group_arn",
				Description: "The Amazon Resource Name (ARN) for the DB subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSubnetGroupArn"),
			},
			{
				Name:        "db_subnet_group_description",
				Description: "Provides the description of the DB subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSubnetGroupDescription"),
			},
			{
				Name:        "db_subnet_group_name",
				Description: "The name of the DB subnet group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSubnetGroupName"),
			},
			{
				Name:        "subnet_group_status",
				Description: "Provides the status of the DB subnet group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_id",
				Description: "Provides the VpcId of the DB subnet group.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_rds_subnet_group_subnets",
				Description: "This data type is used as a response element for the DescribeDBSubnetGroups operation. ",
				Resolver:    fetchRdsSubnetGroupSubnets,
				Columns: []schema.Column{
					{
						Name:        "subnet_group_id",
						Description: "Unique ID of aws_rds_subnet_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "subnet_availability_zone_name",
						Description: "The name of the Availability Zone.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetAvailabilityZone.Name"),
					},
					{
						Name:        "subnet_identifier",
						Description: "The identifier of the subnet.",
						Type:        schema.TypeString,
					},
					{
						Name:        "subnet_outpost_arn",
						Description: "The Amazon Resource Name (ARN) of the Outpost.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SubnetOutpost.Arn"),
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
