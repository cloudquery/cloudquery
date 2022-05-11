package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func RdsDbParameterGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_rds_db_parameter_groups",
		Description:  "Contains the details of an Amazon RDS DB parameter group",
		Resolver:     fetchRdsDbParameterGroups,
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
				Description: "The Amazon Resource Name (ARN) for the DB parameter group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBParameterGroupArn"),
			},
			{
				Name:        "family",
				Description: "The name of the DB parameter group family that this DB parameter group is compatible with.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBParameterGroupFamily"),
			},
			{
				Name:        "name",
				Description: "The name of the DB parameter group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBParameterGroupName"),
			},
			{
				Name:        "description",
				Description: "Provides the customer-specified description for this DB parameter group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "List of tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveRdsDbParameterGroupTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_rds_db_parameters",
				Description: "Database Parameters",
				Resolver:    fetchRdsDbParameterGroupDbParameters,
				IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
				Columns: []schema.Column{
					{
						Name:        "db_parameter_group_cq_id",
						Description: "Unique CloudQuery ID of aws_rds_db_parameter_groups table (FK)",
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
						Name:          "supported_engine_modes",
						Description:   "The valid DB engine modes.",
						Type:          schema.TypeStringArray,
						IgnoreInTests: true,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRdsDbParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().RDS
	var input rds.DescribeDBParameterGroupsInput
	for {
		output, err := svc.DescribeDBParameterGroups(ctx, &input, func(o *rds.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.DBParameterGroups
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func fetchRdsDbParameterGroupDbParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().RDS
	g := parent.Item.(types.DBParameterGroup)
	input := rds.DescribeDBParametersInput{DBParameterGroupName: g.DBParameterGroupName}
	for {
		output, err := svc.DescribeDBParameters(ctx, &input, func(o *rds.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			if client.IsAWSError(err, "DBParameterGroupNotFound") {
				cl.Logger().Debug("received DBParameterGroupNotFound on DescribeDBParameters", "region", cl.Region, "err", err)
				return nil
			}
			return diag.WrapError(err)
		}
		res <- output.Parameters
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func resolveRdsDbParameterGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	g := resource.Item.(types.DBParameterGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().RDS
	out, err := svc.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{ResourceName: g.DBParameterGroupArn}, func(o *rds.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(out.TagList)))
}
