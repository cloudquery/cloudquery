package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ApigatewayAPIKeys() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_api_keys",
		Description:  "A resource that can be distributed to callers for executing Method resources that require an API key.",
		Resolver:     fetchApigatewayApiKeys,
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
				Name:        "created_date",
				Description: "The timestamp when the API Key was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "customer_id",
				Description: "An AWS Marketplace customer identifier , when integrating with the AWS SaaS Marketplace.",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "The description of the API Key.",
				Type:        schema.TypeString,
			},
			{
				Name:        "enabled",
				Description: "Specifies whether the API Key can be used by callers.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "id",
				Description: "The identifier of the API Key.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "last_updated_date",
				Description: "The timestamp when the API Key was last updated.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The name of the API Key.",
				Type:        schema.TypeString,
			},
			{
				Name:        "stage_keys",
				Description: "A list of Stage resources that are associated with the ApiKey resource.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tags",
				Description: "The collection of tags. Each tag element is associated with a given resource.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "value",
				Description: "The value of the API Key.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayApiKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	config := apigateway.GetApiKeysInput{
		IncludeValues: aws.Bool(true),
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for {
		response, err := svc.GetApiKeys(ctx, &config, func(options *apigateway.Options) {
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
