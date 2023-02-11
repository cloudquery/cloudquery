package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ClusterParameters() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_cluster_parameters",
		Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Parameter.html",
		Resolver:    fetchRdsClusterParameters,
		Transform:   transformers.TransformWithStruct(&types.Parameter{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "allowed_values",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AllowedValues"),
			},
			{
				Name:     "apply_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplyMethod"),
			},
			{
				Name:     "apply_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplyType"),
			},
			{
				Name:     "data_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataType"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "is_modifiable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsModifiable"),
			},
			{
				Name:     "minimum_engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MinimumEngineVersion"),
			},
			{
				Name:     "parameter_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ParameterName"),
			},
			{
				Name:     "parameter_value",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ParameterValue"),
			},
			{
				Name:     "source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Source"),
			},
			{
				Name:     "supported_engine_modes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedEngineModes"),
			},
		},
	}
}
