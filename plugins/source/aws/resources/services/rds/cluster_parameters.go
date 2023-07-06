package rds

import (
	"context"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func clusterParameters() *schema.Table {
	tableName := "aws_rds_cluster_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: "https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Parameter.html",
		Resolver:    fetchRdsClusterParameters,
		Transform:   transformers.TransformWithStruct(&types.Parameter{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "allowed_values",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("AllowedValues"),
			},
			{
				Name:     "apply_method",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("ApplyMethod"),
			},
			{
				Name:     "apply_type",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("ApplyType"),
			},
			{
				Name:     "data_type",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("DataType"),
			},
			{
				Name:     "description",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "is_modifiable",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("IsModifiable"),
			},
			{
				Name:     "minimum_engine_version",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("MinimumEngineVersion"),
			},
			{
				Name:     "parameter_name",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("ParameterName"),
			},
			{
				Name:     "parameter_value",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("ParameterValue"),
			},
			{
				Name:     "source",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Source"),
			},
			{
				Name:     "supported_engine_modes",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: schema.PathResolver("SupportedEngineModes"),
			},
		},
	}
}

func fetchRdsClusterParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Rds

	parentEngineVersion := parent.Item.(types.DBEngineVersion)

	if !strings.Contains(aws.ToString(parentEngineVersion.DBParameterGroupFamily), "aurora") {
		return nil
	}

	input := &rds.DescribeEngineDefaultClusterParametersInput{
		DBParameterGroupFamily: parentEngineVersion.DBParameterGroupFamily,
	}

	output, err := svc.DescribeEngineDefaultClusterParameters(ctx, input, func(options *rds.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	if output.EngineDefaults == nil || len(output.EngineDefaults.Parameters) == 0 {
		return nil
	}
	res <- output.EngineDefaults.Parameters
	return nil
}
