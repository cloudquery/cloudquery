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

func ResponseHeaderPolicies() *schema.Table {
	tableName := "aws_cloudfront_response_headers_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListResponseHeadersPolicies.html`,
		Resolver:    fetchCloudfrontResponseHeadersPolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudfront"),
		Transform:   transformers.TransformWithStruct(&types.ResponseHeadersPolicySummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ResponseHeadersPolicy.Id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchCloudfrontResponseHeadersPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	s := cl.Services()
	svc := s.Cloudfront
	var config cloudfront.ListResponseHeadersPoliciesInput
	for {
		response, err := svc.ListResponseHeadersPolicies(ctx, &config, func(options *cloudfront.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		if response.ResponseHeadersPolicyList != nil {
			res <- response.ResponseHeadersPolicyList.Items
		}

		if aws.ToString(response.ResponseHeadersPolicyList.NextMarker) == "" {
			break
		}
		config.Marker = response.ResponseHeadersPolicyList.NextMarker
	}
	return nil
}
