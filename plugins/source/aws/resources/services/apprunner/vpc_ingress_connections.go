package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VpcIngressConnections() *schema.Table {
	tableName := "aws_apprunner_vpc_ingress_connections"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_VpcIngressConnection.html

Notes:
- 'account_id' has been renamed to 'source_account_id' to avoid conflict with the 'account_id' column that indicates what account this was synced from.`,
		Resolver:            fetchApprunnerVpcIngressConnections,
		PreResourceResolver: getVpcIngressConnection,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "apprunner"),
		Transform:           transformers.TransformWithStruct(&types.VpcIngressConnection{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcIngressConnectionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "source_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountId"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveApprunnerTags("VpcIngressConnectionArn"),
			},
		},
	}
}

func fetchApprunnerVpcIngressConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config apprunner.ListVpcIngressConnectionsInput
	svc := meta.(*client.Client).Services().Apprunner
	paginator := apprunner.NewListVpcIngressConnectionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.VpcIngressConnectionSummaryList
	}
	return nil
}

func getVpcIngressConnection(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Apprunner
	asConfig := resource.Item.(types.VpcIngressConnectionSummary)

	describeTaskDefinitionOutput, err := svc.DescribeVpcIngressConnection(ctx, &apprunner.DescribeVpcIngressConnectionInput{VpcIngressConnectionArn: asConfig.VpcIngressConnectionArn})
	if err != nil {
		return err
	}

	resource.Item = describeTaskDefinitionOutput.VpcIngressConnection
	return nil
}
