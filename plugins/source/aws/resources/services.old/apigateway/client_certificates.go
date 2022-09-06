package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource client_certificates --config client_certificates.hcl --output .
func ClientCertificates() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_client_certificates",
		Description:  "Represents a client certificate used to configure client-side SSL authentication while sending requests to the integration endpoint",
		Resolver:     fetchApigatewayClientCertificates,
		Multiplex:    client.ServiceAccountRegionMultiplexer("apigateway"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource",
				Type:        schema.TypeString,
				Resolver:    resolveApigatewayClientCertificateArn,
			},
			{
				Name:        "id",
				Description: "The identifier of the client certificate",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClientCertificateId"),
			},
			{
				Name:        "created_date",
				Description: "The timestamp when the client certificate was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "The description of the client certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "expiration_date",
				Description: "The timestamp when the client certificate will expire",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "pem_encoded_certificate",
				Description: "The PEM-encoded public key of the client certificate, which can be used to configure certificate authentication in the integration endpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The collection of tags",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchApigatewayClientCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigateway.GetClientCertificatesInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for p := apigateway.NewGetClientCertificatesPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayClientCertificateArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	cert := resource.Item.(types.ClientCertificate)
	arn := cl.RegionGlobalARN(client.ApigatewayService, "/clientcertificates", *cert.ClientCertificateId)
	return diag.WrapError(resource.Set(c.Name, arn))
}
