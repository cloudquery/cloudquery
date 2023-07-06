package elbv1

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elbv1/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func LoadBalancers() *schema.Table {
	tableName := "aws_elbv1_load_balancers"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_LoadBalancerDescription.html`,
		Resolver:    fetchElbv1LoadBalancers,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticloadbalancing"),
		Transform:   transformers.TransformWithStruct(&models.ELBv1LoadBalancerWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveLoadBalancerARN(),
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			loadBalancerPolicies(),
		},
	}
}

func fetchElbv1LoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticloadbalancing
	processLoadBalancers := func(loadBalancers []types.LoadBalancerDescription) error {
		tagsCfg := &elbv1.DescribeTagsInput{LoadBalancerNames: make([]string, 0, len(loadBalancers))}
		for _, lb := range loadBalancers {
			tagsCfg.LoadBalancerNames = append(tagsCfg.LoadBalancerNames, *lb.LoadBalancerName)
		}
		tagsResponse, err := svc.DescribeTags(ctx, tagsCfg, func(options *elbv1.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, lb := range loadBalancers {
			loadBalancerAttributes, err := svc.DescribeLoadBalancerAttributes(ctx, &elbv1.DescribeLoadBalancerAttributesInput{LoadBalancerName: lb.LoadBalancerName}, func(options *elbv1.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				if cl.IsNotFoundError(err) {
					continue
				}
				return err
			}

			wrapper := models.ELBv1LoadBalancerWrapper{
				LoadBalancerDescription: lb,
				Tags:                    client.TagsToMap(getTagsByLoadBalancerName(*lb.LoadBalancerName, tagsResponse.TagDescriptions)),
				Attributes:              loadBalancerAttributes.LoadBalancerAttributes,
			}

			res <- wrapper
		}
		return nil
	}
	paginator := elbv1.NewDescribeLoadBalancersPaginator(svc, &elbv1.DescribeLoadBalancersInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *elbv1.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		for i := 0; i < len(page.LoadBalancerDescriptions); i += 20 {
			end := i + 20

			if end > len(page.LoadBalancerDescriptions) {
				end = len(page.LoadBalancerDescriptions)
			}
			loadBalancers := page.LoadBalancerDescriptions[i:end]
			if err := processLoadBalancers(loadBalancers); err != nil {
				return err
			}
		}
	}

	return nil
}

func getTagsByLoadBalancerName(id string, tagsResponse []types.TagDescription) []types.Tag {
	for _, t := range tagsResponse {
		if id == *t.LoadBalancerName {
			return t.Tags
		}
	}
	return nil
}

func resolveLoadBalancerARN() schema.ColumnResolver {
	return client.ResolveARN(client.ElasticLoadBalancingService, func(resource *schema.Resource) ([]string, error) {
		return []string{"loadbalancer", *resource.Item.(models.ELBv1LoadBalancerWrapper).LoadBalancerName}, nil
	})
}
