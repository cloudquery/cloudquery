package redshift

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func clusterParameters() *schema.Table {
	tableName := "aws_redshift_cluster_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_Parameter.html`,
		Resolver:    fetchClusterParameters,
		Transform:   transformers.TransformWithStruct(&types.Parameter{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "cluster_arn",
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.ParentColumnResolver("cluster_arn"),
				Description: `The Amazon Resource Name (ARN) for the resource.`,
				PrimaryKey:  true,
			},
			{
				Name:       "parameter_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ParameterName"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchClusterParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	group := parent.Item.(types.ClusterParameterGroupStatus)
	cl := meta.(*client.Client)
	svc := cl.Services().Redshift

	config := redshift.DescribeClusterParametersInput{
		ParameterGroupName: group.ParameterGroupName,
	}
	paginator := redshift.NewDescribeClusterParametersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *redshift.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Parameters
	}
	return nil
}
