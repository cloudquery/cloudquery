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
		Description:  "Represents a custom domain name as a user-friendly host name of an API (RestApi).",
		Resolver:     fetchApigatewayDomainNames,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name:        "certificate_arn",
				Description: "The reference to an AWS-managed certificate that will be used by edge-optimized endpoint for this domain name. AWS Certificate Manager is the only supported source.",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate_name",
				Description: "The name of the certificate that will be used by edge-optimized endpoint for this domain name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate_upload_date",
				Description: "The timestamp when the certificate that was used by edge-optimized endpoint for this domain name was uploaded.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "distribution_domain_name",
				Description: "The domain name of the Amazon CloudFront distribution associated with this custom domain name for an edge-optimized endpoint. You set up this association when adding a DNS record pointing the custom domain name to this distribution name. For more information about CloudFront distributions, see the Amazon CloudFront documentation (https://aws.amazon.com/documentation/cloudfront/).",
				Type:        schema.TypeString,
			},
			{
				Name:        "distribution_hosted_zone_id",
				Description: "The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized endpoint. The valid value is Z2FDTNDATAQYW2 for all the regions. For more information, see Set up a Regional Custom Domain Name (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-regional-api-custom-domain-create.html) and AWS Regions and Endpoints for API Gateway (https://docs.aws.amazon.com/general/latest/gr/rande.html#apigateway_region).",
				Type:        schema.TypeString,
			},
			{
				Name:        "domain_name",
				Description: "The custom domain name as an API host name, for example, my-api.example.com.",
				Type:        schema.TypeString,
			},
			{
				Name:        "domain_name_status",
				Description: "The status of the DomainName migration. The valid values are AVAILABLE and UPDATING. If the status is UPDATING, the domain cannot be modified further until the existing operation is complete. If it is AVAILABLE, the domain can be updated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "domain_name_status_message",
				Description: "An optional text message containing detailed information about status of the DomainName migration.",
				Type:        schema.TypeString,
			},
			{
				Name:        "endpoint_configuration_types",
				Description: "A list of endpoint types of an API (RestApi) or its custom domain name (DomainName). For an edge-optimized API and its custom domain name, the endpoint type is \"EDGE\". For a regional API and its custom domain name, the endpoint type is REGIONAL. For a private API, the endpoint type is PRIVATE.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointConfiguration.Types"),
			},
			{
				Name:        "endpoint_configuration_vpc_endpoint_ids",
				Description: "A list of VpcEndpointIds of an API (RestApi) against which to create Route53 ALIASes. It is only supported for PRIVATE endpoint type.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointConfiguration.VpcEndpointIds"),
			},
			{
				Name:        "mutual_tls_authentication_truststore_uri",
				Description: "An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example s3://bucket-name/key-name. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version. To update the truststore, you must have permissions to access the S3 object.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MutualTlsAuthentication.TruststoreUri"),
			},
			{
				Name:        "mutual_tls_authentication_truststore_version",
				Description: "The version of the S3 object that contains your truststore. To specify a version, you must have versioning enabled for the S3 bucket.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MutualTlsAuthentication.TruststoreVersion"),
			},
			{
				Name:        "mutual_tls_authentication_truststore_warnings",
				Description: "A list of warnings that API Gateway returns while processing your truststore. Invalid certificates produce warnings. Mutual TLS is still enabled, but some clients might not be able to access your API. To resolve warnings, upload a new truststore to S3, and then update you domain name to use the new version.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("MutualTlsAuthentication.TruststoreWarnings"),
			},
			{
				Name:        "regional_certificate_arn",
				Description: "The reference to an AWS-managed certificate that will be used for validating the regional domain name. AWS Certificate Manager is the only supported source.",
				Type:        schema.TypeString,
			},
			{
				Name:        "regional_certificate_name",
				Description: "The name of the certificate that will be used for validating the regional domain name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "regional_domain_name",
				Description: "The domain name associated with the regional endpoint for this custom domain name. You set up this association by adding a DNS record that points the custom domain name to this regional domain name. The regional domain name is returned by API Gateway when you create a regional endpoint.",
				Type:        schema.TypeString,
			},
			{
				Name:        "regional_hosted_zone_id",
				Description: "The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint. For more information, see Set up a Regional Custom Domain Name (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-regional-api-custom-domain-create.html) and AWS Regions and Endpoints for API Gateway (https://docs.aws.amazon.com/general/latest/gr/rande.html#apigateway_region).",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_policy",
				Description: "The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are TLS_1_0 and TLS_1_2.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The collection of tags. Each tag element is associated with a given resource.",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_apigateway_domain_name_base_path_mappings",
				Description: "Represents the base path that callers of the API must provide as part of the URL after the domain name.",
				Resolver:    fetchApigatewayDomainNameBasePathMappings,
				Columns: []schema.Column{
					{
						Name:        "domain_name_id",
						Description: "Unique ID of aws_apigateway_domain_names table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "base_path",
						Description: "The base path name that callers of the API must provide as part of the URL after the domain name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "rest_api_id",
						Description: "The string identifier of the associated RestApi.",
						Type:        schema.TypeString,
					},
					{
						Name:        "stage",
						Description: "The name of the associated stage.",
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
