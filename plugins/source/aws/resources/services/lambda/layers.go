package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Layers() *schema.Table {
	tableName := "aws_lambda_layers"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_LayersListItem.html`,
		Resolver:    fetchLambdaLayers,
		Transform:   transformers.TransformWithStruct(&types.LayersListItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lambda"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LayerArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			layerVersions(),
		},
	}
}

func fetchLambdaLayers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input lambda.ListLayersInput
	c := meta.(*client.Client)
	svc := c.Services().Lambda
	paginator := lambda.NewListLayersPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *lambda.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		res <- page.Layers
	}
	return nil
}
func fetchLambdaLayerVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(types.LayersListItem)
	cl := meta.(*client.Client)
	svc := cl.Services().Lambda
	config := lambda.ListLayerVersionsInput{
		LayerName: p.LayerName,
	}
	paginator := lambda.NewListLayerVersionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *lambda.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.LayerVersions
	}
	return nil
}
func fetchLambdaLayerVersionPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(types.LayerVersionsListItem)

	pp := parent.Parent.Item.(types.LayersListItem)
	c := meta.(*client.Client)
	svc := c.Services().Lambda

	config := lambda.GetLayerVersionPolicyInput{
		LayerName:     pp.LayerName,
		VersionNumber: p.Version,
	}

	output, err := svc.GetLayerVersionPolicy(ctx, &config, func(options *lambda.Options) {
		options.Region = c.Region
	})
	if err != nil {
		if client.IsAWSError(err, "ResourceNotFoundException") {
			return nil
		}
		return err
	}
	res <- output

	return nil
}
