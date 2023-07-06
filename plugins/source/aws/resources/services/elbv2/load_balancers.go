package elbv2

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func LoadBalancers() *schema.Table {
	tableName := "aws_elbv2_load_balancers"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancer.html`,
		Resolver:    fetchLoadBalancers,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticloadbalancing"),
		Transform:   transformers.TransformWithStruct(&types.LoadBalancer{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveLoadBalancerTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("LoadBalancerArn"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			listeners(),
			loadBalancerAttributes(),
			webACLs(),
		},
	}
}

func fetchLoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config elbv2.DescribeLoadBalancersInput
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticloadbalancingv2
	paginator := elbv2.NewDescribeLoadBalancersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *elbv2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.LoadBalancers
	}
	return nil
}

func resolveLoadBalancerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	region := cl.Region
	svc := cl.Services().Elasticloadbalancingv2
	loadBalancer := resource.Item.(types.LoadBalancer)
	tagsOutput, err := svc.DescribeTags(ctx, &elbv2.DescribeTagsInput{
		ResourceArns: []string{
			*loadBalancer.LoadBalancerArn,
		},
	}, func(o *elbv2.Options) {
		o.Region = region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if len(tagsOutput.TagDescriptions) == 0 {
		return nil
	}
	tags := make(map[string]*string)
	for _, td := range tagsOutput.TagDescriptions {
		for _, s := range td.Tags {
			tags[*s.Key] = s.Value
		}
	}

	return resource.Set(c.Name, tags)
}
