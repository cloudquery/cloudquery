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
		Resolver:     fetchApigatewayApiKeys,
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
				Name: "created_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "customer_id",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "enabled",
				Type: schema.TypeBool,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name: "last_updated_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "stage_keys",
				Type: schema.TypeStringArray,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name: "value",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayApiKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config apigateway.GetApiKeysInput
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
