package elbv2

import (
	"context"

	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func targetGroupTargetHealthDescriptions() *schema.Table {
	tableName := "aws_elbv2_target_group_target_health_descriptions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetHealthDescription.html`,
		Resolver:    fetchTargetGroupTargetHealthDescriptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticloadbalancing"),
		Transform:   transformers.TransformWithStruct(&types.TargetHealthDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "target_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchTargetGroupTargetHealthDescriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticloadbalancingv2
	tg := parent.Item.(types.TargetGroup)
	response, err := svc.DescribeTargetHealth(ctx, &elbv2.DescribeTargetHealthInput{
		TargetGroupArn: tg.TargetGroupArn,
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- response.TargetHealthDescriptions
	return nil
}
