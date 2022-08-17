package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource parameter_groups --config ./gen.hcl --output .
func ParameterGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_elasticache_parameter_groups",
		Description:  "Provides details about Elasticache parameter groups.",
		Resolver:     fetchElasticacheParameterGroups,
		Multiplex:    client.ServiceAccountRegionMultiplexer("elasticache"),
		IgnoreError:  client.IgnoreCommonErrors,
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
				Description: "The ARN (Amazon Resource Name) of the cache parameter group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "cache_parameter_group_family",
				Description: "The name of the cache parameter group family that this cache parameter group is compatible with",
				Type:        schema.TypeString,
			},
			{
				Name:        "cache_parameter_group_name",
				Description: "The name of the cache parameter group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "The description for this cache parameter group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "is_global",
				Description: "Indicates whether the parameter group is associated with a Global datastore",
				Type:        schema.TypeBool,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_elasticache_parameter_group_parameters",
				Description: "Describes an individual setting that controls some aspect of ElastiCache behavior.",
				Resolver:    fetchElasticacheParameterGroupParameters,
				Columns: []schema.Column{
					{
						Name:        "parameter_group_cq_id",
						Description: "Unique CloudQuery ID of aws_elasticache_parameter_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allowed_values",
						Description: "The valid range of values for the parameter.",
						Type:        schema.TypeString,
					},
					{
						Name:        "change_type",
						Description: "Indicates whether a change to the parameter is applied immediately or requires a reboot for the change to be applied",
						Type:        schema.TypeString,
					},
					{
						Name:        "data_type",
						Description: "The valid data type for the parameter.",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "A description of the parameter.",
						Type:        schema.TypeString,
					},
					{
						Name:        "is_modifiable",
						Description: "Indicates whether (true) or not (false) the parameter can be modified",
						Type:        schema.TypeBool,
					},
					{
						Name:        "minimum_engine_version",
						Description: "The earliest cache engine version to which the parameter can apply.",
						Type:        schema.TypeString,
					},
					{
						Name:        "parameter_name",
						Description: "The name of the parameter.",
						Type:        schema.TypeString,
					},
					{
						Name:        "parameter_value",
						Description: "The value of the parameter.",
						Type:        schema.TypeString,
					},
					{
						Name:        "source",
						Description: "The source of the parameter.",
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

func fetchElasticacheParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	awsProviderClient := meta.(*client.Client)
	svc := awsProviderClient.Services().ElastiCache

	var describeCacheParameterGroupsInput elasticache.DescribeCacheParameterGroupsInput

	for {
		describeCacheParameterGroupsOutput, err := svc.DescribeCacheParameterGroups(ctx, &describeCacheParameterGroupsInput)

		if err != nil {
			return diag.WrapError(err)
		}

		res <- describeCacheParameterGroupsOutput.CacheParameterGroups

		if aws.ToString(describeCacheParameterGroupsOutput.Marker) == "" {
			return nil
		}

		describeCacheParameterGroupsInput.Marker = describeCacheParameterGroupsOutput.Marker
	}
}
func fetchElasticacheParameterGroupParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input elasticache.DescribeCacheParametersInput
	parentParameterGroup := parent.Item.(types.CacheParameterGroup)
	input.CacheParameterGroupName = parentParameterGroup.CacheParameterGroupName

	paginator := elasticache.NewDescribeCacheParametersPaginator(meta.(*client.Client).Services().ElastiCache, &input)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- v.Parameters
	}
	return nil
}
