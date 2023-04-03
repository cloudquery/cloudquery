package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func loadBalancerTlsCertificates() *schema.Table {
	tableName := "aws_lightsail_load_balancer_tls_certificates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_LoadBalancerTlsCertificate.html`,
		Resolver:    fetchLightsailLoadBalancerTlsCertificates,
		Transform:   transformers.TransformWithStruct(&types.LoadBalancerTlsCertificate{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "load_balancer_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchLightsailLoadBalancerTlsCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.LoadBalancer)
	input := lightsail.GetLoadBalancerTlsCertificatesInput{
		LoadBalancerName: r.Name,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	response, err := svc.GetLoadBalancerTlsCertificates(ctx, &input)
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- response.TlsCertificates
	return nil
}
