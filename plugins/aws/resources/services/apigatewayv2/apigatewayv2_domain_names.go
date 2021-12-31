package apigatewayv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/cloudquery/cq-provider-aws/client"
	apigatewayv2fix "github.com/cloudquery/cq-provider-aws/resources/forks/apigatewayv2"
)

func Apigatewayv2DomainNames() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigatewayv2_domain_names",
		Description:  "Represents a domain name.",
		Resolver:     fetchApigatewayv2DomainNames,
		Multiplex:    client.ServiceAccountRegionMultiplexer("apigateway"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "domain_name"}},
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
				Name:        "domain_name",
				Description: "The name of the DomainName resource.",
				Type:        schema.TypeString,
			},
			{
				Name:        "api_mapping_selection_expression",
				Description: "The API mapping selection expression.",
				Type:        schema.TypeString,
			},
			{
				Name:        "mutual_tls_authentication_truststore_uri",
				Description: "An Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, s3://bucket-name/key-name. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version. To update the truststore, you must have permissions to access the S3 object.",
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
				Name:        "tags",
				Description: "The collection of tags associated with a domain name.",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_apigatewayv2_domain_name_configurations",
				Description: "The domain name configuration.",
				Resolver:    fetchApigatewayv2DomainNameConfigurations,
				Columns: []schema.Column{
					{
						Name:        "domain_name_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_domain_names table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_gateway_domain_name",
						Description: "A domain name for the API.",
						Type:        schema.TypeString,
					},
					{
						Name:        "certificate_arn",
						Description: "An AWS-managed certificate that will be used by the edge-optimized endpoint for this domain name. AWS Certificate Manager is the only supported source.",
						Type:        schema.TypeString,
					},
					{
						Name:        "certificate_name",
						Description: "The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "certificate_upload_date",
						Description: "The timestamp when the certificate that was used by edge-optimized endpoint for this domain name was uploaded.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "domain_name_status",
						Description: "The status of the domain name migration. The valid values are AVAILABLE and UPDATING. If the status is UPDATING, the domain cannot be modified further until the existing operation is complete. If it is AVAILABLE, the domain can be updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "domain_name_status_message",
						Description: "An optional text message containing detailed information about status of the domain name migration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "endpoint_type",
						Description: "The endpoint type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "hosted_zone_id",
						Description: "The Amazon Route 53 Hosted Zone ID of the endpoint.",
						Type:        schema.TypeString,
					},
					{
						Name:        "security_policy",
						Description: "The Transport Layer Security (TLS) version of the security policy for this domain name. The valid values are TLS_1_0 and TLS_1_2.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_apigatewayv2_domain_name_rest_api_mappings",
				Description: "Represents an API mapping.",
				Resolver:    fetchApigatewayv2DomainNameRestApiMappings,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"domain_name_cq_id", "api_mapping_id"}},
				Columns: []schema.Column{
					{
						Name:        "domain_name_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_domain_names table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API identifier.",
						Type:        schema.TypeString,
					},
					{
						Name:        "stage",
						Description: "The API stage.",
						Type:        schema.TypeString,
					},
					{
						Name:        "api_mapping_id",
						Description: "The API mapping identifier.",
						Type:        schema.TypeString,
					},
					{
						Name:        "api_mapping_key",
						Description: "The API mapping key.",
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

func fetchApigatewayv2DomainNames(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
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

func fetchApigatewayv2DomainNameConfigurations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(types.DomainName)
	if !ok {
		return fmt.Errorf("expected DomainName but got %T", r)
	}
	res <- r.DomainNameConfigurations
	return nil
}

func fetchApigatewayv2DomainNameRestApiMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
