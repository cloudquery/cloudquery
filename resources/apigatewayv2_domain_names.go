package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Apigatewayv2DomainNames() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigatewayv2_domain_names",
		Resolver:     fetchApigatewayv2DomainNames,
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
				Name: "domain_name",
				Type: schema.TypeString,
			},
			{
				Name: "api_mapping_selection_expression",
				Type: schema.TypeString,
			},
			{
				Name:     "mutual_tls_authentication_truststore_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MutualTlsAuthentication.TruststoreUri"),
			},
			{
				Name:     "mutual_tls_authentication_truststore_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MutualTlsAuthentication.TruststoreVersion"),
			},
			{
				Name:     "mutual_tls_authentication_truststore_warnings",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("MutualTlsAuthentication.TruststoreWarnings"),
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_apigatewayv2_domain_name_configurations",
				Resolver: fetchApigatewayv2DomainNameConfigurations,
				Columns: []schema.Column{
					{
						Name:     "domain_name_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "api_gateway_domain_name",
						Type: schema.TypeString,
					},
					{
						Name: "certificate_arn",
						Type: schema.TypeString,
					},
					{
						Name: "certificate_name",
						Type: schema.TypeString,
					},
					{
						Name: "certificate_upload_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "domain_name_status",
						Type: schema.TypeString,
					},
					{
						Name: "domain_name_status_message",
						Type: schema.TypeString,
					},
					{
						Name: "endpoint_type",
						Type: schema.TypeString,
					},
					{
						Name: "hosted_zone_id",
						Type: schema.TypeString,
					},
					{
						Name: "security_policy",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_apigatewayv2_domain_name_api_mappings",
				Resolver: fetchApigatewayv2DomainNameApiMappings,
				Columns: []schema.Column{
					{
						Name:     "domain_name_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "api_id",
						Type: schema.TypeString,
					},
					{
						Name: "stage",
						Type: schema.TypeString,
					},
					{
						Name: "api_mapping_id",
						Type: schema.TypeString,
					},
					{
						Name: "api_mapping_key",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayv2DomainNames(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config apigatewayv2.GetDomainNamesInput
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetDomainNames(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
		})

		if err != nil {
			return err
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func fetchApigatewayv2DomainNameConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.DomainName)
	if !ok {
		return fmt.Errorf("expected DomainName but got %T", r)
	}
	res <- r.DomainNameConfigurations
	return nil
}
func fetchApigatewayv2DomainNameApiMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.DomainName)
	if !ok {
		return fmt.Errorf("expected DomainName but got %T", r)
	}
	config := apigatewayv2.GetApiMappingsInput{
		DomainName: r.DomainName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetApiMappings(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
		})

		if err != nil {
			return err
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
