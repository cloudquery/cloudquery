package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func clusterParameters() *schema.Table {
	tableName := "aws_redshift_cluster_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_Parameter.html`,
		Resolver:    fetchClusterParameters,
		Transform:   client.TransformWithStruct(&types.Parameter{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "cluster_arn",
				Type:        schema.TypeString,
				Resolver:    schema.ParentColumnResolver("cluster_arn"),
				Description: `The Amazon Resource Name (ARN) for the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "parameter_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ParameterName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchClusterParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	group := parent.Item.(types.ClusterParameterGroupStatus)
	c := meta.(*client.Client)
	svc := c.Services().Redshift

	config := redshift.DescribeClusterParametersInput{
		ParameterGroupName: group.ParameterGroupName,
	}
	for {
		response, err := svc.DescribeClusterParameters(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Parameters
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}

	return nil
}
