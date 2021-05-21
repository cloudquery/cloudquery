package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ApigatewayClientCertificates() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_client_certificates",
		Resolver:     fetchApigatewayClientCertificates,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "client_certificate_id",
				Type: schema.TypeString,
			},
			{
				Name: "created_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "expiration_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "pem_encoded_certificate",
				Type: schema.TypeString,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayClientCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
