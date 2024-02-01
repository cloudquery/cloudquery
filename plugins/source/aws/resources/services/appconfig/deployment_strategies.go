package appconfig

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DeploymentStrategies() *schema.Table {
	tableName := "aws_appconfig_deployment_strategies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_DeploymentStrategy.html`,
		Resolver:    fetchDeploymentStrategies,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appconfig"),
		Transform:   transformers.TransformWithStruct(&types.DeploymentStrategy{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveDeploymentStrategiesArn,
				PrimaryKeyComponent: true,
			},
		},
		Relations: schema.Tables{},
	}
}

func fetchDeploymentStrategies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppconfig).Appconfig
	paginator := appconfig.NewListDeploymentStrategiesPaginator(svc, nil)
	for paginator.HasMorePages() {
		resp, err := paginator.NextPage(ctx, func(options *appconfig.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- resp.Items
	}
	return nil
}

// ARN format defined here: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsappconfig.html#awsappconfig-resources-for-iam-policies
// arn:${Partition}:appconfig:${Region}:${Account}:deploymentstrategy/${DeploymentStrategyId}
func resolveDeploymentStrategiesArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	ds := resource.Item.(types.DeploymentStrategy)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.AppconfigService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("deploymentstrategy/%s", aws.ToString(ds.Id)),
	}.String())
}
