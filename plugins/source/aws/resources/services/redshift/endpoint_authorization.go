package redshift

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func endpointAuthorization() *schema.Table {
	tableName := "aws_redshift_endpoint_authorization"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_EndpointAuthorization.html`,
		Resolver:    fetchEndpointAuthorization,
		Transform:   transformers.TransformWithStruct(&types.EndpointAuthorization{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "cluster_arn",
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.ParentColumnResolver("arn"),
				Description: `The Amazon Resource Name (ARN) for the resource.`,
			},
		},
	}
}

func fetchEndpointAuthorization(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(types.Cluster)
	cl := meta.(*client.Client)
	svc := cl.Services().Redshift

	config := redshift.DescribeEndpointAuthorizationInput{
		Account:           &cl.AccountID,
		ClusterIdentifier: cluster.ClusterIdentifier,
		MaxRecords:        aws.Int32(100),
	}
	paginator := redshift.NewDescribeEndpointAuthorizationPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *redshift.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.EndpointAuthorizationList
	}
	return nil
}
