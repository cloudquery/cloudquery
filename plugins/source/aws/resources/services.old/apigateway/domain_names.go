package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource domain_names --config domain_names.hcl --output .
func DomainNames() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_domain_names",
		Description:  "Represents a custom domain name as a user-friendly host name of an API (RestApi)",
		Resolver:     fetchApigatewayDomainNames,
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
				Resolver:    resolveApigatewayDomainNameArn,
			},
			{
				Name:        "certificate_arn",
				Description: "The reference to an AWS-managed certificate that will be used by edge-optimized endpoint for this domain name",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate_name",
				Description: "The name of the certificate that will be used by edge-optimized endpoint for this domain name",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate_upload_date",
				Description: "The timestamp when the certificate that was used by edge-optimized endpoint for this domain name was uploaded",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "distribution_domain_name",
				Description: "The domain name of the Amazon CloudFront distribution associated with this custom domain name for an edge-optimized endpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "distribution_hosted_zone_id",
				Description: "The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized endpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "domain_name",
				Description: "The custom domain name as an API host name, for example, my-api.example.com",
				Type:        schema.TypeString,
			},
			{
				Name:        "domain_name_status",
				Description: "The status of the DomainName migration",
				Type:        schema.TypeString,
			},
			{
				Name:        "domain_name_status_message",
				Description: "An optional text message containing detailed information about status of the DomainName migration",
				Type:        schema.TypeString,
			},
			{
				Name:        "endpoint_configuration_types",
				Description: "A list of endpoint types of an API (RestApi) or its custom domain name (DomainName)",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointConfiguration.Types"),
			},
			{
				Name:        "endpoint_configuration_vpc_endpoint_ids",
				Description: "A list of VpcEndpointIds of an API (RestApi) against which to create Route53 ALIASes",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointConfiguration.VpcEndpointIds"),
			},
			{
				Name:        "mutual_tls_authentication_truststore_uri",
				Description: "An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example s3://bucket-name/key-name",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MutualTlsAuthentication.TruststoreUri"),
			},
			{
				Name:        "mutual_tls_authentication_truststore_version",
				Description: "The version of the S3 object that contains your truststore",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MutualTlsAuthentication.TruststoreVersion"),
			},
			{
				Name:        "mutual_tls_authentication_truststore_warnings",
				Description: "A list of warnings that API Gateway returns while processing your truststore",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("MutualTlsAuthentication.TruststoreWarnings"),
			},
			{
				Name:        "ownership_verification_certificate_arn",
				Description: "The ARN of the public certificate issued by ACM to validate ownership of your custom domain",
				Type:        schema.TypeString,
			},
			{
				Name:        "regional_certificate_arn",
				Description: "The reference to an AWS-managed certificate that will be used for validating the regional domain name",
				Type:        schema.TypeString,
			},
			{
				Name:        "regional_certificate_name",
				Description: "The name of the certificate that will be used for validating the regional domain name",
				Type:        schema.TypeString,
			},
			{
				Name:        "regional_domain_name",
				Description: "The domain name associated with the regional endpoint for this custom domain name",
				Type:        schema.TypeString,
			},
			{
				Name:        "regional_hosted_zone_id",
				Description: "The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_policy",
				Description: "The Transport Layer Security (TLS) version + cipher suite for this DomainName",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The collection of tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_apigateway_domain_name_base_path_mappings",
				Description: "Represents the base path that callers of the API must provide as part of the URL after the domain name",
				Resolver:    fetchApigatewayDomainNameBasePathMappings,
				Columns: []schema.Column{
					{
						Name:        "domain_name_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_domain_names table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayDomainNameBasePathMappingArn,
					},
					{
						Name:        "domain_name",
						Description: "The custom domain name as an API host name",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("DomainName"),
					},
					{
						Name:        "base_path",
						Description: "The base path name that callers of the API must provide as part of the URL after the domain name",
						Type:        schema.TypeString,
					},
					{
						Name:        "rest_api_id",
						Description: "The string identifier of the associated RestApi",
						Type:        schema.TypeString,
					},
					{
						Name:        "stage",
						Description: "The name of the associated stage",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchApigatewayDomainNames(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigateway.GetDomainNamesInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for p := apigateway.NewGetDomainNamesPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayDomainNameArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	domain := resource.Item.(types.DomainName)
	arn := cl.RegionGlobalARN(client.ApigatewayService, domainNameIDPart, *domain.DomainName)
	return diag.WrapError(resource.Set(c.Name, arn))
}
func fetchApigatewayDomainNameBasePathMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.DomainName)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetBasePathMappingsInput{DomainName: r.DomainName}
	for p := apigateway.NewGetBasePathMappingsPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayDomainNameBasePathMappingArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	domain := resource.Parent.Item.(types.DomainName)
	mapping := resource.Item.(types.BasePathMapping)
	arn := cl.RegionGlobalARN(client.ApigatewayService, domainNameIDPart, *domain.DomainName, "basepathmappings", *mapping.BasePath)
	return diag.WrapError(resource.Set(c.Name, arn))
}
