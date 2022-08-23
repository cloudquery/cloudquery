package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	apigatewayv2fix "github.com/cloudquery/cloudquery/plugins/source/aws/resources/forks/apigatewayv2"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource domain_names --config gen.hcl --output .
func DomainNames() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigatewayv2_domain_names",
		Description:  "Represents a domain name",
		Resolver:     fetchApigatewayv2DomainNames,
		Multiplex:    client.ServiceAccountRegionMultiplexer("apigateway"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Description: "The Amazon Resource Name (ARN) for the resource",
				Type:        schema.TypeString,
				Resolver:    resolveApigatewayv2domainNameArn,
			},
			{
				Name:        "domain_name",
				Description: "The name of the DomainName resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "api_mapping_selection_expression",
				Description: "The API mapping selection expression",
				Type:        schema.TypeString,
			},
			{
				Name:        "mutual_tls_authentication_truststore_uri",
				Description: "An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, s3://bucket-name/key-name",
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
				Name:        "tags",
				Description: "The collection of tags associated with a domain name",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_apigatewayv2_domain_name_configurations",
				Description: "The domain name configuration",
				Resolver:    schema.PathTableResolver("DomainNameConfigurations"),
				Columns: []schema.Column{
					{
						Name:        "domain_name_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_domain_names table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_gateway_domain_name",
						Description: "A domain name for the API",
						Type:        schema.TypeString,
					},
					{
						Name:        "certificate_arn",
						Description: "An AWS-managed certificate that will be used by the edge-optimized endpoint for this domain name",
						Type:        schema.TypeString,
					},
					{
						Name:        "certificate_name",
						Description: "The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name",
						Type:        schema.TypeString,
					},
					{
						Name:        "certificate_upload_date",
						Description: "The timestamp when the certificate that was used by edge-optimized endpoint for this domain name was uploaded",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "domain_name_status",
						Description: "The status of the domain name migration",
						Type:        schema.TypeString,
					},
					{
						Name:        "domain_name_status_message",
						Description: "An optional text message containing detailed information about status of the domain name migration",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoint_type",
						Description: "The endpoint type",
						Type:        schema.TypeString,
					},
					{
						Name:        "hosted_zone_id",
						Description: "The Amazon Route 53 Hosted Zone ID of the endpoint",
						Type:        schema.TypeString,
					},
					{
						Name:        "ownership_verification_certificate_arn",
						Description: "The ARN of the public certificate issued by ACM to validate ownership of your custom domain",
						Type:        schema.TypeString,
					},
					{
						Name:        "security_policy",
						Description: "The Transport Layer Security (TLS) version of the security policy for this domain name",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_apigatewayv2_domain_name_rest_api_mappings",
				Description: "Represents an API mapping",
				Resolver:    fetchApigatewayv2DomainNameRestApiMappings,
				Columns: []schema.Column{
					{
						Name:        "domain_name_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_domain_names table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayv2domainNameRestAPIMappingArn,
					},
					{
						Name:        "api_id",
						Description: "The API identifier",
						Type:        schema.TypeString,
					},
					{
						Name:        "stage",
						Description: "The API stage",
						Type:        schema.TypeString,
					},
					{
						Name:        "api_mapping_id",
						Description: "The API mapping identifier",
						Type:        schema.TypeString,
					},
					{
						Name:        "api_mapping_key",
						Description: "The API mapping key",
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

func fetchApigatewayv2DomainNames(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigatewayv2.GetDomainNamesInput
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetDomainNames(ctx, &config, func(options *apigatewayv2.Options) {
			options.Region = c.Region
			// NOTE: Swapping OperationDeserializer until this is fixed: https://github.com/aws/aws-sdk-go-v2/issues/1282
			options.APIOptions = append(options.APIOptions, apigatewayv2fix.SwapGetDomainNamesOperationDeserializer)
		})

		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveApigatewayv2domainNameArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	d := resource.Item.(types.DomainName)
	arn := cl.RegionGlobalARN(client.ApigatewayService, domainNamesIDPart, *d.DomainName)
	return diag.WrapError(resource.Set(c.Name, arn))
}
func fetchApigatewayv2DomainNameRestApiMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.DomainName)
	config := apigatewayv2.GetApiMappingsInput{
		DomainName: r.DomainName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetApiMappings(ctx, &config)

		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveApigatewayv2domainNameRestAPIMappingArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	d := resource.Parent.Item.(types.DomainName)
	m := resource.Item.(types.ApiMapping)
	arn := cl.RegionGlobalARN(client.ApigatewayService, domainNamesIDPart, *d.DomainName, "apimappings", *m.ApiMappingId)
	return diag.WrapError(resource.Set(c.Name, arn))
}
