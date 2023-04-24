package elbv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/v2/transformers"
)

func listenerCertificates() *schema.Table {
	tableName := "aws_elbv2_listener_certificates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Certificate.html`,
		Resolver:    fetchListenerCertificates,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticloadbalancing"),
		Transform:   transformers.TransformWithStruct(&types.Certificate{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "listener_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchListenerCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	region := c.Region
	svc := c.Services().Elasticloadbalancingv2
	listener := parent.Item.(types.Listener)
	config := elbv2.DescribeListenerCertificatesInput{ListenerArn: listener.ListenerArn}
	// No paginator available
	for {
		response, err := svc.DescribeListenerCertificates(ctx, &config, func(options *elbv2.Options) {
			options.Region = region
		})
		if err != nil {
			return err
		}
		res <- response.Certificates
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
