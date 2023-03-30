package elbv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TargetGroups() *schema.Table {
	tableName := "aws_elbv2_target_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetGroup.html`,
		Resolver:    fetchTargetGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticloadbalancing"),
		Transform:   transformers.TransformWithStruct(&types.TargetGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveTargetGroupTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			targetGroupTargetHealthDescriptions(),
		},
	}
}

func fetchTargetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config elbv2.DescribeTargetGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().Elasticloadbalancingv2
	for {
		response, err := svc.DescribeTargetGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.TargetGroups
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}

func resolveTargetGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().Elasticloadbalancingv2
	targetGroup := resource.Item.(types.TargetGroup)
	tagsOutput, err := svc.DescribeTags(ctx, &elbv2.DescribeTagsInput{
		ResourceArns: []string{
			*targetGroup.TargetGroupArn,
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
	for _, s := range tagsOutput.TagDescriptions[0].Tags {
		tags[*s.Key] = s.Value
	}
	return resource.Set(c.Name, tags)
}
