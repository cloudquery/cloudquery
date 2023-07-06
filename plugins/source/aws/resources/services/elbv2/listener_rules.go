package elbv2

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func listenerRules() *schema.Table {
	tableName := "aws_elbv2_listener_rules"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Rule.html`,
		Resolver:    fetchListenerRules,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticloadbalancing"),
		Transform:   transformers.TransformWithStruct(&types.Rule{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "listener_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("RuleArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchListenerRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	region := cl.Region
	svc := cl.Services().Elasticloadbalancingv2
	listener := parent.Item.(types.Listener)
	config := elbv2.DescribeRulesInput{ListenerArn: listener.ListenerArn}
	// no paginator available
	for {
		response, err := svc.DescribeRules(ctx, &config, func(options *elbv2.Options) {
			options.Region = region
		})
		if err != nil {
			return err
		}
		res <- response.Rules
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
