package lightsail

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func loadBalancerTlsCertificates() *schema.Table {
	tableName := "aws_lightsail_load_balancer_tls_certificates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_LoadBalancerTlsCertificate.html`,
		Resolver:    fetchLightsailLoadBalancerTlsCertificates,
		Transform:   transformers.TransformWithStruct(&types.LoadBalancerTlsCertificate{}, transformers.WithPrimaryKeyComponents("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "load_balancer_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
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
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLightsail).Lightsail
	response, err := svc.GetLoadBalancerTlsCertificates(ctx, &input, func(options *lightsail.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- response.TlsCertificates
	return nil
}
