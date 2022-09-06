package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource api_keys --config api_keys.hcl --output .
func APIKeys() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_api_keys",
		Description:  "A resource that can be distributed to callers for executing Method resources that require an API key",
		Resolver:     fetchApigatewayApiKeys,
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
				Resolver:    resolveApigatewayAPIKeyArn,
			},
			{
				Name:        "created_date",
				Description: "The timestamp when the API Key was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "customer_id",
				Description: "An AWS Marketplace customer identifier , when integrating with the AWS SaaS Marketplace",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "The description of the API Key",
				Type:        schema.TypeString,
			},
			{
				Name:        "enabled",
				Description: "Specifies whether the API Key can be used by callers",
				Type:        schema.TypeBool,
			},
			{
				Name:        "id",
				Description: "The identifier of the API Key",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_updated_date",
				Description: "The timestamp when the API Key was last updated",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The name of the API Key",
				Type:        schema.TypeString,
			},
			{
				Name:        "stage_keys",
				Description: "A list of Stage resources that are associated with the ApiKey resource",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tags",
				Description: "The collection of tags",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "value",
				Description: "The value of the API Key",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchApigatewayApiKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	config := apigateway.GetApiKeysInput{
		IncludeValues: aws.Bool(true),
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	p := apigateway.NewGetApiKeysPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayAPIKeyArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	ak := resource.Item.(types.ApiKey)
	arn := cl.RegionGlobalARN(client.ApigatewayService, "/apikeys", *ak.Id)
	return diag.WrapError(resource.Set(c.Name, arn))
}
