package rds

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Instances() *schema.Table {
	tableName := "aws_rds_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBInstance.html`,
		Resolver:    fetchRdsInstances,
		Transform:   transformers.TransformWithStruct(&types.DBInstance{}, transformers.WithSkipFields("TagList")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("DBInstanceArn"),
				PrimaryKey: true,
			},
			{
				Name:     "processor_features",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRdsInstanceProcessorFeatures,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRdsInstanceTags,
			},
		},
	}
}

func fetchRdsInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config rds.DescribeDBInstancesInput
	cl := meta.(*client.Client)
	svc := cl.Services().Rds
	paginator := rds.NewDescribeDBInstancesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DBInstances
	}
	return nil
}

func resolveRdsInstanceProcessorFeatures(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DBInstance)
	processorFeatures := map[string]*string{}
	for _, t := range r.ProcessorFeatures {
		processorFeatures[*t.Name] = t.Value
	}
	return resource.Set(c.Name, processorFeatures)
}

func resolveRdsInstanceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return resource.Set(c.Name, client.TagsToMap(resource.Item.(types.DBInstance).TagList))
}
