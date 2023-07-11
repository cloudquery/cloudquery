package cloudfront

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Functions() *schema.Table {
	tableName := "aws_cloudfront_functions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_DescribeFunction.html`,
		Resolver:            fetchFunctions,
		PreResourceResolver: getFunction,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cloudfront"),
		Transform:           transformers.TransformWithStruct(&cloudfront.DescribeFunctionOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:       "stage",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("FunctionSummary.FunctionMetadata.Stage"),
				PrimaryKey: true,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("FunctionSummary.FunctionMetadata.FunctionARN"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudfront.ListFunctionsInput
	cl := meta.(*client.Client)
	s := cl.Services()
	svc := s.Cloudfront
	for {
		response, err := svc.ListFunctions(ctx, &config, func(options *cloudfront.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		if response.FunctionList != nil && len(response.FunctionList.Items) > 0 {
			res <- response.FunctionList.Items
		}

		if aws.ToString(response.FunctionList.NextMarker) == "" {
			break
		}
		config.Marker = response.FunctionList.NextMarker
	}
	return nil
}
func getFunction(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudfront

	f := resource.Item.(types.FunctionSummary)

	function, err := svc.DescribeFunction(ctx, &cloudfront.DescribeFunctionInput{
		Name:  f.Name,
		Stage: f.FunctionMetadata.Stage,
	}, func(options *cloudfront.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = function
	return nil
}
