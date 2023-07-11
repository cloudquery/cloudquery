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

func OriginRequestPolicies() *schema.Table {
	tableName := "aws_cloudfront_origin_request_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListOriginRequestPolicies.html`,
		Resolver:    fetchCloudfrontOriginRequestPolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudfront"),
		Transform:   transformers.TransformWithStruct(&types.OriginRequestPolicySummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("OriginRequestPolicy.Id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchCloudfrontOriginRequestPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	s := cl.Services()
	svc := s.Cloudfront
	var config cloudfront.ListOriginRequestPoliciesInput
	for {
		response, err := svc.ListOriginRequestPolicies(ctx, &config, func(options *cloudfront.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		if response.OriginRequestPolicyList != nil && len(response.OriginRequestPolicyList.Items) > 0 {
			res <- response.OriginRequestPolicyList.Items
		}

		if aws.ToString(response.OriginRequestPolicyList.NextMarker) == "" {
			break
		}
		config.Marker = response.OriginRequestPolicyList.NextMarker
	}
	return nil
}
