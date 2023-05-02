package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lambda/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func functionAliases() *schema.Table {
	tableName := "aws_lambda_function_aliases"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/lambda/latest/dg/API_AliasConfiguration.html`,
		Resolver:            fetchLambdaFunctionAliases,
		PreResourceResolver: getFunctionAliasURLConfig,
		Transform:           transformers.TransformWithStruct(&models.AliasWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "lambda"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "function_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AliasArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchLambdaFunctionAliases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	c := meta.(*client.Client)
	svc := c.Services().Lambda
	config := lambda.ListAliasesInput{
		FunctionName: p.Configuration.FunctionName,
	}
	paginator := lambda.NewListAliasesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *lambda.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- page.Aliases
	}
	return nil
}

func getFunctionAliasURLConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Lambda
	alias := resource.Item.(types.AliasConfiguration)
	p := resource.Parent.Item.(*lambda.GetFunctionOutput)

	urlConfig, err := svc.GetFunctionUrlConfig(ctx, &lambda.GetFunctionUrlConfigInput{
		FunctionName: p.Configuration.FunctionName,
		Qualifier:    alias.Name,
	}, func(options *lambda.Options) {
		options.Region = c.Region
	})
	if err != nil && !c.IsNotFoundError(err) {
		return err
	}

	resource.Item = &models.AliasWrapper{AliasConfiguration: &alias, UrlConfig: urlConfig}
	return nil
}
