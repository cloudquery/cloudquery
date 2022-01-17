package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ApigatewayClientCertificates() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_client_certificates",
		Description:  "Represents a client certificate used to configure client-side SSL authentication while sending requests to the integration endpoint.",
		Resolver:     fetchApigatewayClientCertificates,
		Multiplex:    client.ServiceAccountRegionMultiplexer("apigateway"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
					return []string{"/clientcertificates", *resource.Item.(types.ClientCertificate).ClientCertificateId}, nil
				}),
			},
			{
				Name:        "id",
				Description: "The identifier of the client certificate.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ClientCertificateId"),
			},
			{
				Name:        "created_date",
				Description: "The timestamp when the client certificate was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "The description of the client certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "expiration_date",
				Description: "The timestamp when the client certificate will expire.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "pem_encoded_certificate",
				Description: "The PEM-encoded public key of the client certificate, which can be used to configure certificate authentication in the integration endpoint .",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The collection of tags. Each tag element is associated with a given resource.",
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
	for {
		response, err := svc.GetClientCertificates(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
