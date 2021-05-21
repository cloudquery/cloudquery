package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ApigatewayDomainNames() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_domain_names",
		Resolver:     fetchApigatewayDomainNames,
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
				Name: "distribution_domain_name",
				Type: schema.TypeString,
			},
			{
				Name: "distribution_hosted_zone_id",
				Type: schema.TypeString,
			},
			{
				Name: "domain_name",
				Type: schema.TypeString,
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
				Name:     "endpoint_configuration_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EndpointConfiguration.Types"),
			},
			{
				Name:     "endpoint_configuration_vpc_endpoint_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EndpointConfiguration.VpcEndpointIds"),
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
				Name: "regional_certificate_arn",
				Type: schema.TypeString,
			},
			{
				Name: "regional_certificate_name",
				Type: schema.TypeString,
			},
			{
				Name: "regional_domain_name",
				Type: schema.TypeString,
			},
			{
				Name: "regional_hosted_zone_id",
				Type: schema.TypeString,
			},
			{
				Name: "security_policy",
				Type: schema.TypeString,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_apigateway_domain_name_base_path_mappings",
				Resolver: fetchApigatewayDomainNameBasePathMappings,
				Columns: []schema.Column{
					{
						Name:     "domain_name_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "base_path",
						Type: schema.TypeString,
					},
					{
						Name: "rest_api_id",
						Type: schema.TypeString,
					},
					{
						Name: "stage",
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
func fetchApigatewayDomainNames(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config apigateway.GetDomainNamesInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for {
		response, err := svc.GetDomainNames(ctx, &config, func(options *apigateway.Options) {
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
func fetchApigatewayDomainNameBasePathMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.DomainName)
	if !ok {
		return fmt.Errorf("expected DomainName but got %T", r)
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetBasePathMappingsInput{DomainName: r.DomainName}
	for {
		response, err := svc.GetBasePathMappings(ctx, &config, func(options *apigateway.Options) {
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
