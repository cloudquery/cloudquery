package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func RdsClusterParameterGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_rds_cluster_parameter_groups",
		Description:  "Contains the details of an Amazon RDS DB cluster parameter group",
		Resolver:     fetchRdsClusterParameterGroups,
		Multiplex:    client.ServiceAccountRegionMultiplexer("rds"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Description: "The Amazon Resource Name (ARN) for the DB cluster parameter group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBClusterParameterGroupArn"),
			},
			{
				Name:        "name",
				Description: "The name of the DB cluster parameter group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBClusterParameterGroupName"),
			},
			{
				Name:        "family",
				Description: "The name of the DB parameter group family that this DB cluster parameter group is compatible with.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBParameterGroupFamily"),
			},
			{
				Name:        "description",
				Description: "Provides the customer-specified description for this DB cluster parameter group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "List of tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveRdsClusterParameterGroupTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_rds_cluster_parameters",
				Description: "This data type is used as a request parameter in the ModifyDBParameterGroup and ResetDBParameterGroup actions",
				Resolver:    fetchRdsClusterParameterGroupDbParameters,
				IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
				Columns: []schema.Column{
					{
						Name:        "cluster_parameter_group_cq_id",
						Description: "Unique CloudQuery ID of aws_rds_cluster_parameter_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allowed_values",
						Description: "Specifies the valid range of values for the parameter.",
						Type:        schema.TypeString,
					},
					{
						Name:        "apply_method",
						Description: "Indicates when to apply parameter updates.",
						Type:        schema.TypeString,
					},
					{
						Name:        "apply_type",
						Description: "Specifies the engine specific parameters type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "data_type",
						Description: "Specifies the valid data type for the parameter.",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "Provides a description of the parameter.",
						Type:        schema.TypeString,
					},
					{
						Name:        "is_modifiable",
						Description: "Indicates whether (true) or not (false) the parameter can be modified",
						Type:        schema.TypeBool,
					},
					{
						Name:        "minimum_engine_version",
						Description: "The earliest engine version to which the parameter can apply.",
						Type:        schema.TypeString,
					},
					{
						Name:        "parameter_name",
						Description: "Specifies the name of the parameter.",
						Type:        schema.TypeString,
					},
					{
						Name:        "parameter_value",
						Description: "Specifies the value of the parameter.",
						Type:        schema.TypeString,
					},
					{
						Name:        "source",
						Description: "Indicates the source of the parameter value.",
						Type:        schema.TypeString,
					},
					{
						Name:        "supported_engine_modes",
						Description: "The valid DB engine modes.",
						Type:        schema.TypeStringArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRdsClusterParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().RDS
	var input rds.DescribeDBClusterParameterGroupsInput
	for {
		output, err := svc.DescribeDBClusterParameterGroups(ctx, &input, func(o *rds.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.DBClusterParameterGroups
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func fetchRdsClusterParameterGroupDbParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().RDS
	g := parent.Item.(types.DBClusterParameterGroup)
	input := rds.DescribeDBClusterParametersInput{DBClusterParameterGroupName: g.DBClusterParameterGroupName}
	for {
		output, err := svc.DescribeDBClusterParameters(ctx, &input, func(o *rds.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Parameters
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func resolveRdsClusterParameterGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	g := resource.Item.(types.DBClusterParameterGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().RDS
	out, err := svc.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{ResourceName: g.DBClusterParameterGroupArn}, func(o *rds.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	tags := make(map[string]string, len(out.TagList))
	for _, t := range out.TagList {
		tags[aws.ToString(t.Key)] = aws.ToString(t.Value)
	}
	return resource.Set(c.Name, tags)
}
