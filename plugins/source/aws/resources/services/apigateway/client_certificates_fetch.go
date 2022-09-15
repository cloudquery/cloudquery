package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchApigatewayClientCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigateway.GetClientCertificatesInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for p := apigateway.NewGetClientCertificatesPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayClientCertificateArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	cert := resource.Item.(types.ClientCertificate)
	arn := cl.RegionGlobalARN(client.ApigatewayService, "/clientcertificates", *cert.ClientCertificateId)
	return resource.Set(c.Name, arn)
}
