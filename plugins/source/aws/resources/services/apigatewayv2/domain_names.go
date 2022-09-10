package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	apigatewayv2fix "github.com/cloudquery/cloudquery/plugins/source/aws/resources/forks/apigatewayv2"
	"github.com/cloudquery/plugin-sdk/schema"
)

const domainNamesIDPart = "domainnames"

func Apigatewayv2DomainNames() *schema.Table {
	return &schema.Table{
		Name:          "aws_apigatewayv2_domain_names",
		Description:   "Represents a domain name.",
		Resolver:      fetchApigatewayv2DomainNames,
		Multiplex:     client.ServiceAccountRegionMultiplexer("apigateway"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "region",
				Description:     "The AWS Region of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "domain_name",
				Description:     "The name of the DomainName resource.",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
					return []string{domainNamesIDPart, *resource.Item.(types.DomainName).DomainName}, nil
				}),
			},
			{
				Name:        "api_mapping_selection_expression",
				Description: "The API mapping selection expression.",
				Type:        schema.TypeString,
			},
			{
				Name:     "mutual_tls_authentication",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MutualTlsAuthentication"),
			},
			{
				Name:        "tags",
				Description: "The collection of tags associated with a domain name.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "configurations",
				Description: "The domain name configuration.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("DomainNameConfigurations"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_apigatewayv2_domain_name_rest_api_mappings",
				Description: "Represents an API mapping.",
				Resolver:    fetchApigatewayv2DomainNameRestApiMappings,
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
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.ApiMapping)
							p := resource.Parent.Item.(types.DomainName)
							return []string{domainNamesIDPart, *p.DomainName, "apimappings", *r.ApiMappingId}, nil
						}),
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
