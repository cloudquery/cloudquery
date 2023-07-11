package elbv2

import (
	"context"
	"errors"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func webACLs() *schema.Table {
	tableName := "aws_elbv2_load_balancer_web_acls"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_GetWebACLForResource.html`,
		Resolver:    resolveLoadBalancerWebACL,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "waf-regional"),
		Transform:   transformers.TransformWithStruct(&wafv2types.WebACL{}, transformers.WithPrimaryKeys("ARN")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "load_balancer_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func resolveLoadBalancerWebACL(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(types.LoadBalancer)
	// only application load balancer can have web acl arn
	if p.Type != types.LoadBalancerTypeEnumApplication {
		return nil
	}
	cl := meta.(*client.Client)
	wafClient := cl.Services().Wafv2
	input := wafv2.GetWebACLForResourceInput{ResourceArn: p.LoadBalancerArn}
	response, err := wafClient.GetWebACLForResource(ctx, &input, func(options *wafv2.Options) {}, func(options *wafv2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		var exc *wafv2types.WAFNonexistentItemException
		if errors.As(err, &exc) {
			if exc.ErrorCode() == "WAFNonexistentItemException" {
				return nil
			}
		}

		return err
	}
	res <- response.WebACL
	return nil
}
